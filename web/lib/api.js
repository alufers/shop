import create from "axios";

export class APIError extends Error {
  constructor(code, message) {
    super(message);
    this.code = code;
  }
}

const api = create.create({
  baseURL: "http://localhost:8080"
});

// api.interceptors.response.use(function(response) {
//   if (response.code > 399) {
//     throw new APIError(response.data.errors[0]);
//   }
//   return response;
// });

export default api;
