import React, { useState } from "react";
import useSWR, { mutate } from "swr";
import { useForm } from "react-hook-form";

import { Select, Label, Text } from "../../../../components/form";
import Button from "../../../../components/button";
import Modal from "../../../../components/modal";
import { ReactComponent as Loader } from "../../../../assets/three-dots.svg";

import { User, Course, Department } from "../../../../interfaces";
import useSWRPost from "../../../../hooks/useSWRPost";

import styles from "../users.module.css";

interface Payload {
  name: string;
  email: string;
  password: string;
  role: User["role"];
  rollNo?: string;
  admissionTime?: string;
  courseEndTime?: string;
  designation?: string;
  salutation?: string;
  courseID?: string;
  departmentID?: string;
}

interface WrapperProps {
  handleSubmit: any;
  onSubmit(v: Payload): void;
  register: any;
  setIsOpen(v: boolean): void;
  isValidating: boolean;
  values: Payload;
}

const WrapperStep: React.FC<WrapperProps> = ({
  children,
  values: v,
  handleSubmit,
  onSubmit,
  register,
  setIsOpen,
  isValidating,
}) => {
  return (
    <form onSubmit={handleSubmit(onSubmit)} className={styles.addForm}>
      <h2>Create New User</h2>
      <input type="hidden" name="name" defaultValue={v.name} ref={register({ required: true })} />
      <input type="hidden" name="email" defaultValue={v.email} ref={register({ required: true })} />
      <input
        type="hidden"
        name="password"
        defaultValue={v.password}
        ref={register({ required: true })}
      />
      <input type="hidden" name="role" defaultValue={v.role} ref={register({ required: true })} />
      {children}
      <Button
        type="button"
        onClick={() => setIsOpen(false)}
        style={{ background: "var(--color-danger)" }}
        disabled={isValidating}
      >
        Cancel
      </Button>
      <Button style={{ background: "var(--color-success)" }} disabled={isValidating}>
        {isValidating ? <Loader style={{ height: "1rem" }} /> : `Submit`}
      </Button>
    </form>
  );
};

interface StepProps {
  values: Payload;
  onSubmit(v: Payload): void;
  setIsOpen(v: boolean): void;
  isValidating: boolean;
}

const StepStaff = ({ values, onSubmit, setIsOpen, isValidating }: StepProps) => {
  const { register, handleSubmit } = useForm<Payload>();
  return (
    <WrapperStep
      values={values}
      register={register}
      handleSubmit={handleSubmit}
      isValidating={isValidating}
      setIsOpen={setIsOpen}
      onSubmit={onSubmit}
    >
      <Label value="Designation">
        <Text
          name="designation"
          placeholder="Ex. Registrar"
          inpRef={register({ required: true })}
        />
      </Label>
    </WrapperStep>
  );
};

const StepFaculty = ({ values, onSubmit, setIsOpen, isValidating }: StepProps) => {
  const { data } = useSWR<{ status: string; departments: Department[] }>("/api/departments");

  const { register, handleSubmit } = useForm<Payload>();
  return (
    <WrapperStep
      values={values}
      register={register}
      handleSubmit={handleSubmit}
      isValidating={isValidating}
      setIsOpen={setIsOpen}
      onSubmit={onSubmit}
    >
      <Label value="Salutation">
        <Text name="salutation" placeholder="Dr." inpRef={register()} />
      </Label>
      <Label value="Department">
        <Select name="departmentID" inpRef={register({ required: true })}>
          {data?.departments.map((d) => (
            <option value={d.id} key={d.id}>
              {d.name}
            </option>
          ))}
        </Select>
      </Label>
    </WrapperStep>
  );
};

const StepStudent = ({ values, onSubmit, setIsOpen, isValidating }: StepProps) => {
  const { data } = useSWR<{ status: string; courses: Course[] }>("/api/courses");

  const { register, handleSubmit } = useForm<Payload>();
  return (
    <WrapperStep
      values={values}
      register={register}
      handleSubmit={handleSubmit}
      isValidating={isValidating}
      setIsOpen={setIsOpen}
      onSubmit={onSubmit}
    >
      <Label value="Roll No.">
        <Text name="rollNo" placeholder="20XXIMX-0XX" inpRef={register({ required: true })} />
      </Label>
      <Label value="Admission Time">
        <Text type="date" name="admissionTime" inpRef={register({ required: true })} />
      </Label>
      <Label value="Course End Time">
        <Text type="date" name="courseEndTime" inpRef={register({ required: true })} />
      </Label>
      <Label value="Course">
        <Select name="courseID" inpRef={register({ required: true })}>
          {data?.courses.map((c) => (
            <option value={c.id} key={c.id}>
              {c.name}
            </option>
          ))}
        </Select>
      </Label>
    </WrapperStep>
  );
};

const StepOne = ({
  setIsOpen,
  isValidating,
  finalSubmit,
}: {
  setIsOpen(v: boolean): void;
  isValidating: boolean;
  finalSubmit(v: Payload): void;
}) => {
  const [role, setRole] = useState<User["role"]>();

  const { register, handleSubmit, getValues } = useForm<Payload>();

  if (role === "student")
    return (
      <StepStudent
        values={getValues()}
        onSubmit={finalSubmit}
        setIsOpen={setIsOpen}
        isValidating={isValidating}
      />
    );
  else if (role === "faculty")
    return (
      <StepFaculty
        values={getValues()}
        onSubmit={finalSubmit}
        setIsOpen={setIsOpen}
        isValidating={isValidating}
      />
    );
  else if (role === "staff")
    return (
      <StepStaff
        values={getValues()}
        onSubmit={finalSubmit}
        setIsOpen={setIsOpen}
        isValidating={isValidating}
      />
    );

  const onSubmit = (values: Payload) => {
    if (values.role === "admin" || values.role === "misc") finalSubmit(values);
    else setRole(values.role);
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} className={styles.addForm}>
      <h2>Create New User</h2>
      <Label value="Name *">
        <Text
          name="name"
          placeholder="John Doe"
          inpRef={register({
            required: true,
          })}
        />
      </Label>
      <Label value="Email *">
        <Text
          name="email"
          type="email"
          placeholder="example@iiitm.ac.in"
          inpRef={register({
            required: true,
          })}
        />
      </Label>
      <Label value="Password *">
        <Text
          name="password"
          type="password"
          autoComplete="new-password"
          placeholder="s3cRetP4$s"
          inpRef={register({
            required: true,
          })}
        />
      </Label>
      <Label value="Role *">
        <Select
          name="role"
          inpRef={register({
            required: true,
          })}
        >
          <option value="student">Student</option>
          <option value="faculty">Faculty</option>
          <option value="staff">Staff</option>
          <option value="misc">Misc</option>
          <option value="admin">Admin</option>
        </Select>
      </Label>
      <Button
        type="button"
        onClick={() => setIsOpen(false)}
        style={{ background: "var(--color-danger)" }}
        disabled={isValidating}
      >
        Cancel
      </Button>
      <Button style={{ background: "var(--color-brand)" }} disabled={isValidating}>
        {isValidating ? <Loader style={{ height: "1rem" }} /> : `Next`}
      </Button>
    </form>
  );
};

const AddUser = () => {
  const [isOpen, setIsOpen] = useState<boolean>(false);
  const [role, setRole] = useState<User["role"]>();
  const [createUser, { isValidating }] = useSWRPost<string>("/api/user", {
    onSuccess: (res) => {
      if (res.status === "success") {
        mutate(`/api/users?role=${role}`);
        setIsOpen(false);
      } else console.log(res);
    },
    onError: console.log,
  });

  const onSubmit = (values: Payload) => {
    const p = {
      ...values,
      admissionTime: values.admissionTime ? new Date(values.admissionTime) : null,
      courseEndTime: values.courseEndTime ? new Date(values.courseEndTime) : null,
      courseID: Number(values.courseID),
      departmentID: Number(values.departmentID),
    };
    createUser(JSON.stringify(p));
    setRole(values.role);
  };
  return (
    <>
      <Button
        type="button"
        style={{ background: "var(--color-brand)", width: "max-content" }}
        onClick={() => setIsOpen(true)}
      >
        Add New User
      </Button>
      <Modal isOpen={isOpen} toggleModal={setIsOpen}>
        <StepOne setIsOpen={setIsOpen} isValidating={isValidating} finalSubmit={onSubmit} />
      </Modal>
    </>
  );
};

export default AddUser;
