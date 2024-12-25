import axiosInstance from './axiosConfig';

export const getAllOrders = async () => {
  try {
    const response = await axiosInstance.get('/orders/');
    return response.data;
  } catch (error) {
    console.error('Error fetching orders:', error);
    throw error;
  }
};

export const getOrderById = async (id) => {
  try {
    const response = await axiosInstance.get(`/orders/${id}`);
    return response.data;
  } catch (error) {
    console.error('Error fetching order by ID:', error);
    throw error;
  }
};

export const createOrder = async (orderData) => {
  try {
    const response = await axiosInstance.post('/orders/', orderData);
    return response.data;
  } catch (error) {
    console.error('Error creating order:', error);
    throw error;
  }
};

export const updateOrder = async (id, orderData) => {
  try {
    const response = await axiosInstance.put(`/orders/${id}`, orderData);
    return response.data;
  } catch (error) {
    console.error('Error updating order:', error);
    throw error;
  }
};

export const updateOrderStatus = async (id, status) => {
  try {
    const response = await axiosInstance.put(`/orders/status/${id}`, { status });
    return response.data;
  } catch (error) {
    console.error('Error updating order status:', error);
    throw error;
  }
};

export const deleteOrder = async (id) => {
  try {
    const response = await axiosInstance.delete(`/orders/${id}`);
    return response.data;
  } catch (error) {
    console.error('Error deleting order:', error);
    throw error;
  }
};

export const searchOrders = async (query) => {
  try {
    const response = await axiosInstance.get('/orders/search', { params: { q: query } });
    return response.data;
  } catch (error) {
    console.error('Error searching orders:', error);
    throw error;
  }
};

export const getOrderByUserId = async (userId) => {
  try {
    const response = await axiosInstance.get(`/orders/user/${userId}`);
    return response.data;
  } catch (error) {
    console.error('Error fetching orders by user ID:', error);
    throw error;
  }
};

export const cancelOrder = async (id) => {
  try {
    return response.data;
  } catch (error) {
    console.error('Error canceling order:', error);
    throw error;
  }
};

export const updateOrderStatusByOrderId = async (orderId, status) => {
  try {
    const response = await axiosInstance.put(`/status/${orderId}`, { status });
    return response.data;
  } catch (error) {
    console.error('Error updating order status by order ID:', error);
    throw error;
  }
};

export const addOrderDetail = async (orderDetailData) => {
  try {
    const response = await axiosInstance.post('/orders/detail', orderDetailData);
    return response.data;
  } catch (error) {
    console.error('Error adding order detail:', error);
    throw error;
  }
};

export const updateOrderDetail = async (orderDetailData) => {
  try {
    const response = await axiosInstance.put('/orders/detail', orderDetailData);
    return response.data;
  } catch (error) {
    console.error('Error updating order detail:', error);
    throw error;
  }
};

export const deleteOrderDetail = async (orderDetailId) => {
  try {
    const response = await axiosInstance.delete(`/orders/detail/${orderDetailId}`);
    return response.data;
  } catch (error) {
    console.error('Error deleting order detail:', error);
    throw error;
  }
};

export const getOrderDetailsByOrderId = async (orderId) => {
  try {
    const response = await axiosInstance.get(`/orders/detail/order/${orderId}`);
    return response.data;
  } catch (error) {
    console.error('Error fetching order details by order ID:', error);
    throw error;
  }
};

export const getAllOrderDetails = async () => {
  try {
    const response = await axiosInstance.get('/orders/detail');
    return response.data;
  } catch (error) {
    console.error('Error fetching all order details:', error);
    throw error;
  }
};