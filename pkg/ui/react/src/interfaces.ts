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

  // social handles
  linkedin?: string;
  facebook?: string;
  github?: string;
  twitter?: string;

  edges: {
    // student
    Course?: Course;
    // faculty
    Department?: Department;
  };
}

export interface Course {
  id: number;
  name: string;
  code: string;
  semesters: number;
}

export interface Department {
  id: number;
  name: string;
}

export interface ExtClient {
  client: Client;
  payload: ClientPayload;
}

export interface Client {
  id: number;
  name: string;
  clientID: string;
  secret: string;

  edges: {
    User?: User;
  };
}

export interface ClientPayload {
  allowed_cors_origins: string[];
  client_id: string;
  client_name: string;
  created_at: string;
  grant_types: string[];
  logo_uri: string;
  metadata: Record<string, string>;
  owner: string;
  redirect_uris: string[];
  response_types: string[];
  scope: string;
  subject_type: string;
  token_endpoint_auth_method: string;
  updated_at: string;
  userinfo_signed_response_alg: string;
}
