export interface User {
  name: string;
  email: string;
  role: "student" | "faculty" | "staff" | "admin";
  photo?: string
  altEmail?: string
}
