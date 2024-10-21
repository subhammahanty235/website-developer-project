import axios from "axios";
const cloudinaryUpload = async (file) => {
    const cloudName = 'dbnqqpobe'; // Replace with your Cloudinary cloud name
    const uploadPreset = 'bozeifzkyeueuhsggwebsitecreator'; // Replace with your Cloudinary unsigned upload preset

    const formData = new FormData();
    formData.append('file', file);
    formData.append('upload_preset', uploadPreset);

    try {
        const response = await axios.post(`https://api.cloudinary.com/v1_1/${cloudName}/image/upload`, formData, {
            headers: {
                'Content-Type': 'multipart/form-data',
            },
        });

        return response.data.secure_url; // Return the URL of the uploaded image
    } catch (error) {
        console.error('Error uploading image to Cloudinary:', error);
        throw error;
    }
};

export default cloudinaryUpload