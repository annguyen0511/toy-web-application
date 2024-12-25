<template>
  <div>
    <a-breadcrumb style="margin: 16px 0">
      <a-breadcrumb-item>Trang chủ</a-breadcrumb-item>
      <a-breadcrumb-item>Quản lý đơn hàng</a-breadcrumb-item>
    </a-breadcrumb>
    <div class="actions-bar">
      <a-input-search
        placeholder="Tìm kiếm đơn hàng"
        enter-button
        @search="handleSearch"
        style="max-width: 300px;"
      />
    </div>
    <a-table
      :columns="columns"
      :dataSource="orders"
      :pagination="{ pageSize: 10 }"
      rowKey="order_id"
    >
      <template #bodyCell="{ column, record }">
        <span v-if="column.key === 'actions' && record.order_status !== 'Đã hủy'">
          <a @click="showEditModal(record)" style="margin-right: 8px;">Cập nhật trạng thái</a>
          <a @click="showViewModal(record)" style="margin-right: 8px;">Xem chi tiết</a>
        </span>
        <span v-else-if="column.key === 'order_date'">
          {{ formatDate(record.order_date) }}
        </span>
        <span v-else-if="column.key === 'total_amount'">
          {{ formatCurrency(record.total_amount) }}
        </span>
        <span v-else>{{ record[column.dataIndex] || '-' }}</span>
      </template>
    </a-table>
    <a-modal
      v-model:visible="isModalVisible"
      title="Cập nhật trạng thái đơn hàng"
      @ok="handleOk"
      @cancel="handleCancel"
      width="600px"
    >
      <a-form :model="formData" ref="orderForm" layout="vertical">
        <a-form-item label="Trạng thái" name="order_status" style="margin-bottom: 10px;">
          <a-select v-model:value="formData.order_status" placeholder="Chọn trạng thái">
            <a-select-option value="Từ chối">Từ chối</a-select-option>
            <a-select-option value="Vận chuyển">Vận chuyển</a-select-option>
            <a-select-option value="Phê duyệt">Phê duyệt</a-select-option>
            <a-select-option value="Đợi xác nhận">Đợi xác nhận</a-select-option>
          </a-select>
        </a-form-item>
      </a-form>
    </a-modal>
    <a-modal
      v-model:visible="isViewModalVisible"
      title="Chi tiết đơn hàng"
      @cancel="handleViewCancel"
      width="800px"
    >
      <div v-if="selectedOrderBooks.length">
        <div v-for="book in selectedOrderBooks" :key="book.order_detail_id" class="book-detail">
          <p>Sản phẩm ID: {{ book.product_id }}</p>
          <p>Số lượng: {{ book.quantity }}</p>
          <p>Giá: {{ formatCurrency(book.price) }}</p>
        </div>
      </div>
      <div v-else>
        <p>Không có thông tin sách để hiển thị.</p>
      </div>
    </a-modal>
  </div>
</template>

<script>
import { getAllOrders, updateOrderStatus, getOrderDetailsByOrderId } from '@/apis/ordersApi';
import { message } from 'ant-design-vue';
import moment from 'moment';

export default {
  name: 'OrderManagement',
  data() {
    return {
      orders: [],
      columns: [
        {
          title: 'Mã đơn hàng',
          dataIndex: 'order_id',
          key: 'order_id',
        },
        {
          title: 'Mã người dùng',
          dataIndex: 'user_id',
          key: 'user_id',
        },
        {
          title: 'Ngày đặt hàng',
          dataIndex: 'order_date',
          key: 'order_date',
        },
        {
          title: 'Tổng giá',
          dataIndex: 'total_amount',
          key: 'total_amount',
        },
        {
          title: 'Trạng thái',
          dataIndex: 'order_status',
          key: 'order_status',
        },
        {
          title: 'Hành động',
          key: 'actions',
        },
      ],
      isModalVisible: false,
      isViewModalVisible: false,
      formData: {
        order_id: null,
        order_status: '',
      },
      selectedOrderBooks: [],
    };
  },
  async created() {
    this.fetchOrders();
  },
  methods: {
    async fetchOrders() {
      try {
        const data = await getAllOrders();
        this.orders = data;
      } catch (error) {
        message.error(error.message || 'Có lỗi xảy ra khi tải đơn hàng!');
      }
    },
    async handleSearch(query) {
      if (!query) {
        this.fetchOrders();
        return;
      }
      this.orders = this.orders.filter(order => 
        String(order.order_id).includes(query) || 
        String(order.user_id).includes(query)
      );
    },
    showEditModal(record) {
      this.formData = { order_id: record.order_id, order_status: record.order_status };
      this.isModalVisible = true;
    },
    async handleOk() {
      try {
        await updateOrderStatus(this.formData.order_id, this.formData.order_status);
        message.success('Cập nhật trạng thái đơn hàng thành công!');
        this.isModalVisible = false;
        this.fetchOrders();
      } catch (error) {
        message.error(error.message || 'Có lỗi xảy ra!');
      }
    },
    handleCancel() {
      this.isModalVisible = false;
    },
    showViewModal(record) {
      this.selectedOrderBooks = record.books || [];
      this.isViewModalVisible = true;
      this.fetchOrderDetails(record.order_id);
    },
    async fetchOrderDetails(orderId) {
      try {
        const data = await getOrderDetailsByOrderId(orderId);
        this.selectedOrderBooks = data || [];
      } catch (error) {
        message.error(error.message || 'Có lỗi xảy ra khi tải chi tiết đơn hàng!');
      }
    },
    handleViewCancel() {
      this.isViewModalVisible = false;
    },
    formatDate(date) {
      return moment(date).format('DD/MM/YYYY');
    },
    formatCurrency(value) {
      return new Intl.NumberFormat('vi-VN', { style: 'currency', currency: 'VND' }).format(value);
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
.book-detail {
  display: flex;
  margin-bottom: 16px;
}
.book-image {
  width: 100px;
  height: auto;
  margin-right: 16px;
}
.book-info {
  flex: 1;
}
</style> 