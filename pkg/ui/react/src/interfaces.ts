export interface User {
  id: number;
  name: string;
  email: string;
  role: "student" | "faculty" | "staff" | "admin";
  photo?: string;
  altEmail?: string;
  phone?: string;
  // student
  rollNo?: string;
  admissionTime?: string;
  courseEndTime?: string;

  edges: UserEdges;
}

export interface UserEdges {
  // student
  Course?: {
    id: number;
    name: string;
    code: string;
    semesters: number;
  };
  // faculty
  Department?: {
    id: number;
    name: string;
  };
}
