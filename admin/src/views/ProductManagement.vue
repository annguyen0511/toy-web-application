<template>
  <div>
    <a-breadcrumb style="margin: 16px 0">
      <a-breadcrumb-item>Trang chủ</a-breadcrumb-item>
      <a-breadcrumb-item>Quản lý sản phẩm</a-breadcrumb-item>
    </a-breadcrumb>
    <div class="actions-bar">
      <a-input-search
        placeholder="Tìm kiếm sản phẩm"
        enter-button
        @search="handleSearch"
        style="max-width: 300px;"
      />
      <a-button type="primary" @click="showCreateModal">Thêm sản phẩm</a-button>
    </div>
    <a-table
      :columns="columns"
      :dataSource="products"
      :pagination="{ pageSize: 10 }"
      rowKey="id"
    >
      <template #bodyCell="{ column, record }">
        <span v-if="column.key === 'actions'">
          <a @click="showEditModal(record)">Sửa</a>
          <a-divider type="vertical" />
          <a-popconfirm
            title="Bạn có chắc chắn muốn xóa sản phẩm này không?"
            ok-text="Có"
            cancel-text="Không"
            @confirm="handleDelete(record.product_id)"
          >
            <a>Xóa</a>
          </a-popconfirm>
        </span>
        <img v-else-if="column.key === 'image_url'" :src="record.image_url || placeholderImage" alt="Image" style="width: 50px; height: 50px;" />
        <span v-else-if="column.key === 'price'">
          {{ formatCurrency(record.price) }}
        </span>
        <span v-else>{{ record[column.dataIndex] || '-' }}</span>
      </template>
    </a-table>
    <a-modal
      v-model:visible="isModalVisible"
      title="Sản phẩm"
      @ok="handleOk"
      @cancel="handleCancel"
    >
      <a-form :model="formData" :rules="rules" ref="productForm" layout="vertical">
        <a-form-item label="Tên sản phẩm" name="product_name" style="margin-bottom: 10px;">
          <a-input v-model:value="formData.product_name" />
        </a-form-item>
        <a-form-item label="Mô tả" name="description" style="margin-bottom: 10px;">
          <a-textarea v-model:value="formData.description" />
        </a-form-item>
        <a-form-item label="Giá" name="price" style="margin-bottom: 10px;">
          <a-input-number v-model:value="formData.price" style="width: 100%;" />
        </a-form-item>
        <a-form-item label="Số lượng tồn kho" name="stock" style="margin-bottom: 10px;">
          <a-input-number v-model:value="formData.stock" style="width: 100%;" />
        </a-form-item>
        <a-form-item label="Danh mục" name="category_id" style="margin-bottom: 10px;">
          <a-select v-model:value="formData.category_id" placeholder="Chọn danh mục">
            <a-select-option v-for="category in categories" :key="category.category_id" :value="category.category_id">
              {{ category.category_name }}
            </a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item label="Hình ảnh" name="image_url" style="margin-bottom: 10px;">
          <a-upload
            name="file"
            :customRequest="handleCustomRequest"
            list-type="picture-card"
            :show-upload-list="false"
          >
            <div>
              <a-icon type="plus" />
              <div style="margin-top: 8px">Upload</div>
            </div>
          </a-upload>
          <img v-if="formData.image_url" :src="formData.image_url" alt="Image" style="width: 100px; margin-top: 8px;" />
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script>
import { getAllProducts, createProduct, updateProduct, deleteProduct, searchProducts, updateOrderStatusByOrderId } from '@/apis/productApi';
import { getAllCategories } from '@/apis/categoriesApi';
import { message } from 'ant-design-vue';
import _ from 'lodash';

export default {
  name: 'ProductManagement',
  data() {
    return {
      products: [],
      categories: [],
      columns: [
        {
          title: 'Hình ảnh',
          key: 'image_url',
        },
        {
          title: 'Tên sản phẩm',
          dataIndex: 'product_name',
          key: 'product_name',
        },
        {
          title: 'Mô tả',
          dataIndex: 'description',
          key: 'description',
        },
        {
          title: 'Giá',
          dataIndex: 'price',
          key: 'price',
        },
        {
          title: 'Số lượng tồn kho',
          dataIndex: 'stock',
          key: 'stock',
        },
        {
          title: 'Danh mục',
          dataIndex: 'category_name',
          key: 'category_name',
        },
        {
          title: 'Hành động',
          key: 'actions',
        },
      ],
      isModalVisible: false,
      formData: {
        id: null,
        product_name: '',
        description: '',
        price: 0,
        stock: 0,
        category_id: '',
        image_url: '',
      },
      isEditing: false,
      placeholderImage: 'https://www.svgrepo.com/show/508699/landscape-placeholder.svg', // Optional: Set a placeholder image path
      rules: {
        product_name: [{ required: true, message: 'Vui lòng nhập tên sản phẩm!' }],
        description: [{ required: true, message: 'Vui lòng nhập mô tả!' }],
        price: [{ required: true, message: 'Vui lòng nhập giá!' }],
        stock: [{ required: true, message: 'Vui lòng nhập số lượng tồn kho!' }],
        category_id: [{ required: true, message: 'Vui lòng chọn danh mục!' }],
        image_url: [{ required: true, message: 'Vui lòng tải lên hình ảnh!' }],
      },
    };
  },
  async created() {
    this.fetchProducts();
    this.fetchCategories();
  },
  methods: {
    async fetchProducts() {
     try {
       const data = await getAllProducts();
       console.log(data); // Log the data to check its structure
       await this.fetchCategories(); // Ensure categories are fetched before processing products
       this.products = data.map(product => {
         const category = this.categories.find(cat => Number(cat.category_id) === Number(product.category_id));
         return {
           ...product,
           category_name: category ? category.category_name : 'Không xác định',
         };
       });
     } catch (error) {
       console.log(error);
       message.error(error.message || 'Có lỗi xảy ra khi tải sản phẩm!');
     }
   },
    async fetchCategories() {
      try {
        const data = await getAllCategories();
        this.categories = data;
      } catch (error) {
        console.log(error);
        message.error(error.message || 'Có lỗi xảy ra khi tải danh mục!');
      }
    },
    async handleSearch(query) {
      try {
        const data = await searchProducts(query);
        this.products = data;
      } catch (error) {
        message.error(error.message || 'Có lỗi xảy ra khi tìm kiếm sản phẩm!');
      }
    },
    showCreateModal() {
      this.isEditing = false;
      this.formData = {
        id: null,
        product_name: '',
        description: '',
        price: 0,
        stock: 0,
        category_id: '',
        image_url: '',
      };
      this.isModalVisible = true;
    },
    showEditModal(record) {
      this.isEditing = true;
      this.formData = { ...record };
      this.fetchCategories().then(() => {
        const category = this.categories.find(cat => cat.category_id === record.category_id);
        this.formData.category_id = category ? category.category_id : null;
      });
      this.isModalVisible = true;
    },
    async handleOk() {
      this.$refs.productForm.validate().then(async () => {
        try {
          const payload = _.pick(this.formData, [
            'product_name',
            'description',
            'price',
            'stock',
            'category_id',
            'image_url',
          ]);

          if (this.isEditing) {
            await updateProduct(this.formData.product_id, { ...payload, product_id: this.formData.product_id });
            message.success('Cập nhật sản phẩm thành công!');
          } else {
            await createProduct(payload);
            message.success('Thêm sản phẩm thành công!');
          }
          this.isModalVisible = false;
          this.fetchProducts();
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
    async handleDelete(id) {
      try {
        await deleteProduct(id);
        message.success('Xóa sản phẩm thành công!');
        this.fetchProducts();
      } catch (error) {
        message.error(error.message || 'Có lỗi xảy ra!');
      }
    },
    async handleCustomRequest({ file, onSuccess, onError }) {
      try {
        const formData = new FormData();
        formData.append('image', file);
        const response = await fetch(`https://api.imgbb.com/1/upload?key=e64a49ca517de7491f78d8edf586515a`, {
          method: 'POST',
          body: formData,
        });
        const data = await response.json();
        
        if (data.success) {
          this.formData.image_url = data.data.url;
          onSuccess(data, file);
          message.success(`${file.name} file uploaded successfully`);
        } else {
          throw new Error('Upload failed');
        }
      } catch (error) {
        onError(error);
        message.error(`${file.name} file upload failed.`);
      }
    },
    formatCurrency(value) {
      return new Intl.NumberFormat('vi-VN', { style: 'currency', currency: 'VND' }).format(value);
    },
    async updateOrderStatus(orderId, status) {
      try {
        await updateOrderStatusByOrderId(orderId, { status });
        message.success('Cập nhật trạng thái đơn hàng thành công!');
      } catch (error) {
        message.error(error.message || 'Có lỗi xảy ra khi cập nhật trạng thái đơn hàng!');
      }
    },
    async handleUpdateOrderStatus(orderId) {
      const newStatus = 'Shipped';
      await this.updateOrderStatus(orderId, newStatus);
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