import axios from "axios"

export const template_instance = axios.create({
   // baseURL : "http://localhost:8001/template/",
    baseURL :  process.env.VUE_APP_TEMPLATE_SERVER_URL + "/template/",
});

export const auth_instance = axios.create({
   // baseURL : "http://localhost:8000/auth/",
     baseURL :  process.env.VUE_APP_AUTH_SERVER_URL + "/auth/",
});


export default {
    signUpUser(data){
        return auth_instance.post( 'sign-up', data)
    },
    signInUser(data){
        return auth_instance.post( 'sign-in', data)
    },
    createTemplate(data, jwt){
        let config = {
            headers: {
                'Authorization': 'Bearer ' + jwt,
            },
        }
        return template_instance.post( 'create', data, config)
    },
    getAllTemplates(jwt) {
        let config = {
            headers: {
                'Authorization': 'Bearer ' + jwt,
            },
        }
        return template_instance.get(``, config)
    },
    getTemplateById(selectedItem, jwt) {
        let config = {
            headers: {
                'Authorization': 'Bearer ' + jwt,
            },
        }
        console.log(selectedItem);
        return template_instance.get(`${selectedItem}`, config)
    },
    updateTemplate(selectedItem, data, jwt) {
        let config = {
            headers: {
                'Authorization': 'Bearer ' + jwt,
            },
        }
        return template_instance.patch(`${selectedItem}`, data, config)
    },
    deleteTemplate(selectedItem, jwt) {
        console.log(selectedItem)
        let config = {
            headers: {
                'Authorization': 'Bearer ' + jwt,
            },
        }
        return template_instance.delete(`${selectedItem}`, config)
    }
}
