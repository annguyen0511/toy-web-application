import axiosInstance from './axiosConfig';

export const getAllUsers = async () => {
  try {
    const response = await axiosInstance.get('/users/');
    return response.data;
  } catch (error) {
    console.error('Error fetching users:', error);
    throw error;
  }
};

export const getUserProfile = async () => {
  try {
    const response = await axiosInstance.get('/users/profile');
    return response.data;
  } catch (error) {
    console.error('Error fetching user profile:', error);
    throw error;
  }
};

export const searchUserByEmail = async (email) => {
  try {
    const response = await axiosInstance.get(`/users/searchByEmail`, { params: { email } });
    return response.data;
  } catch (error) {
    console.error('Error searching user by email:', error);
    throw error;
  }
};

export const createUser = async (userData) => {
  try {
    const response = await axiosInstance.post('/users/register', userData);
    return response.data;
  } catch (error) {
    console.error('Error creating user:', error);
    throw error;
  }
};

export const updateUser = async (id, userData) => {
  try {
    const response = await axiosInstance.put(`/users/${id}`, userData);
    return response.data;
  } catch (error) {
    console.error('Error updating user:', error);
    throw error;
  }
};

export const deleteUser = async (id) => {
  try {
    const response = await axiosInstance.delete(`/users/${id}`);
    return response.data;
  } catch (error) {
    console.error('Error deleting user:', error);
    throw error;
  }
};

export const changePassword = async (userId, newPassword) => {
  try {
    const response = await axiosInstance.post('/users/change-password', { userId, newPassword });
    return response.data;
  } catch (error) {
    console.error('Error changing password:', error);
    throw error;
  }
};

export const resetPassword = async (email, newPassword) => {
  try {
    const response = await axiosInstance.post('/users/reset-password', { email, newPassword });
    return response.data;
  } catch (error) {
    console.error('Error resetting password:', error);
    throw error;
  }
};

export const sendResetPasswordEmail = async (email) => {
  try {
    const response = await axiosInstance.post('/users/send-reset-password-email', { email });
    return response.data;
  } catch (error) {
    console.error('Error sending reset password email:', error);
    throw error;
  }
};
