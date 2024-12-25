import axiosInstance from './axiosConfig';

export const getAllUsers = async () => {
  try {
    const response = await axiosInstance.get('/user');
    return response.data;
  } catch (error) {
    console.error('Error fetching users:', error);
    throw error;
  }
};

export const getUserProfile = async () => {
  try {
    const response = await axiosInstance.get('/user/profile');
    return response.data;
  } catch (error) {
    console.error('Error fetching user profile:', error);
    throw error;
  }
};

export const searchUserByEmail = async (email) => {
  try {
    const response = await axiosInstance.get(`/user/searchByEmail`, { params: { email } });
    return response.data;
  } catch (error) {
    console.error('Error searching user by email:', error);
    throw error;
  }
};

export const createUser = async (userData) => {
  try {
    const response = await axiosInstance.post('/user', userData);
    return response.data;
  } catch (error) {
    console.error('Error creating user:', error);
    throw error;
  }
};

export const updateUser = async (id, userData) => {
  try {
    const response = await axiosInstance.put(`/user/${id}`, userData);
    return response.data;
  } catch (error) {
    console.error('Error updating user:', error);
    throw error;
  }
};

export const deleteUser = async (id) => {
  try {
    const response = await axiosInstance.delete(`/user/${id}`);
    return response.data;
  } catch (error) {
    console.error('Error deleting user:', error);
    throw error;
  }
};

export const registerUser = async (userData) => {
  try {
    const response = await axiosInstance.post('/users/register', userData);
    return response.data;
  } catch (error) {
    console.error('Error registering user:', error);
    throw error;
  }
};

export const loginUser = async (credentials) => {
  try {
    const response = await axiosInstance.post('/user/login', credentials);
    return response.data;
  } catch (error) {
    console.error('Error logging in user:', error);
    throw error;
  }
};

export const changePassword = async (userId, passwordData) => {
  try {
    const response = await axiosInstance.post(`/user/change-password`, { userId, ...passwordData });
    return response.data;
  } catch (error) {
    console.error('Error changing password:', error);
    throw error;
  }
};

export const resetPassword = async (email) => {
  try {
    const response = await axiosInstance.post('/user/reset-password', { email });
    return response.data;
  } catch (error) {
    console.error('Error resetting password:', error);
    throw error;
  }
};

export const sendResetPasswordEmail = async (email) => {
  try {
    const response = await axiosInstance.post('/user/send-reset-password-email', { email });
    return response.data;
  } catch (error) {
    console.error('Error sending reset password email:', error);
    throw error;
  }
};
