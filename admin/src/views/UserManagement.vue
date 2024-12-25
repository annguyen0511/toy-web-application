<template>
  <div>
    <a-breadcrumb style="margin: 16px 0">
      <a-breadcrumb-item>Trang chủ</a-breadcrumb-item>
      <a-breadcrumb-item>Quản lý người dùng</a-breadcrumb-item>
    </a-breadcrumb>
    <div class="actions-bar">
      <a-input-search
        placeholder="Tìm kiếm người dùng"
        enter-button
        @search="handleSearch"
        style="max-width: 300px;"
      />
      <a-button type="primary" @click="showCreateModal">Thêm người dùng</a-button>
    </div>
    <a-table
      :columns="columns"
      :dataSource="users"
      :pagination="{ pageSize: 10 }"
      rowKey="user_id"
    >
      <template #bodyCell="{ column, record }">
        <span v-if="column.key === 'actions'">
          <a @click="showEditModal(record)">Sửa</a>
          <a-divider type="vertical" />
          <a-popconfirm
            title="Bạn có chắc chắn muốn xóa người dùng này không?"
            ok-text="Có"
            cancel-text="Không"
            @confirm="handleDelete(record.user_id)"
          >
            <a>Xóa</a>
          </a-popconfirm>
        </span>
        <span v-else>{{ record[column.dataIndex] || '-' }}</span>
      </template>
    </a-table>
    <a-modal
      v-model:visible="isModalVisible"
      title="Người dùng"
      @ok="handleOk"
      @cancel="handleCancel"
    >
      <a-form :model="formData" :rules="rules" ref="userForm" layout="vertical">
        <a-form-item label="Họ và tên" name="full_name" style="margin-bottom: 10px;">
          <a-input v-model:value="formData.full_name" />
        </a-form-item>
        <a-form-item label="Email" name="email" style="margin-bottom: 10px;">
          <a-input v-model:value="formData.email" />
        </a-form-item>
        <a-form-item label="Số điện thoại" name="phone_number" style="margin-bottom: 10px;">
          <a-input v-model:value="formData.phone_number" />
        </a-form-item>
        <a-form-item label="Địa chỉ" name="address" style="margin-bottom: 10px;">
          <a-input v-model:value="formData.address" />
        </a-form-item>
        <a-form-item label="Vai trò" name="role" style="margin-bottom: 10px;">
          <a-select v-model:value="formData.role" placeholder="Chọn vai trò">
            <a-select-option value="User">Người dùng</a-select-option>
            <a-select-option value="isAdmin">Quản trị viên</a-select-option>
          </a-select>
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script>
import { getAllUsers, createUser, updateUser, deleteUser, searchUserByEmail } from '@/apis/userApi';
import { message } from 'ant-design-vue';

export default {
  name: 'UserManagement',
  data() {
    return {
      users: [],
      columns: [
        {
          title: 'Họ và tên',
          dataIndex: 'full_name',
          key: 'full_name',
        },
        {
          title: 'Email',
          dataIndex: 'email',
          key: 'email',
        },
        {
          title: 'Số điện thoại',
          dataIndex: 'phone_number',
          key: 'phone_number',
        },
        {
          title: 'Địa chỉ',
          dataIndex: 'address',
          key: 'address',
        },
        {
          title: 'Vai trò',
          dataIndex: 'role',
          key: 'role',
        },
        {
          title: 'Hành động',
          key: 'actions',
        },
      ],
      isModalVisible: false,
      formData: {
        user_id: null,
        full_name: '',
        email: '',
        phone_number: '',
        address: '',
        role: '',
      },
      isEditing: false,
      rules: {
        full_name: [{ required: true, message: 'Vui lòng nhập họ và tên!' }],
        email: [{ required: true, message: 'Vui lòng nhập email!' }],
        phone_number: [{ required: true, message: 'Vui lòng nhập số điện thoại!' }],
        address: [{ required: true, message: 'Vui lòng nhập địa chỉ!' }],
        role: [{ required: true, message: 'Vui lòng chọn vai trò!' }],
      },
    };
  },
  async created() {
    this.fetchUsers();
  },
  methods: {
    async fetchUsers() {
      try {
        const data = await getAllUsers();
        this.users = data;
      } catch (error) {
        message.error(error.message || 'Có lỗi xảy ra khi tải người dùng!');
      }
    },
    async handleSearch(query) {
      try {
        const data = await searchUserByEmail(query);
        this.users = data;
      } catch (error) {
        message.error(error.message || 'Có lỗi xảy ra khi tìm kiếm người dùng!');
      }
    },
    showCreateModal() {
      this.isEditing = false;
      this.formData = {
        user_id: null,
        full_name: '',
        email: '',
        phone_number: '',
        address: '',
        role: '',
      };
      this.isModalVisible = true;
    },
    showEditModal(record) {
      this.isEditing = true;
      this.formData = { ...record }; 
      this.isModalVisible = true;
    },
    async handleOk() {
      this.$refs.userForm.validate().then(async () => {
        try {
          if (this.isEditing) {
            await updateUser(this.formData.user_id, this.formData);
            message.success('Cập nhật người dùng thành công!');
          } else {
            await createUser(this.formData);
            message.success('Thêm người dùng thành công!');
          }
          this.isModalVisible = false;
          this.fetchUsers();
        } catch (error) {
          message.error(error.message || 'Có lỗi xảy ra!');
        }
      }).catch(() => {
        message.error('Vui lòng điền đầy đủ thông tin!');
      });
    },
    handleCancel() {
      this.isModalVisible = false;
    },
    async handleDelete(user_id) {
      try {
        await deleteUser(user_id);
        message.success('Xóa người dùng thành công!');
        this.fetchUsers();
      } catch (error) {
        message.error(error.message || 'Có lỗi xảy ra!');
      }
    },
  },
};
</script>

<style scoped>
.actions-bar {
  display: flex;
  justify-content: space-between;
  margin-bottom: 16px;
}
</style> 