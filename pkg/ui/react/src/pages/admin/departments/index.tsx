import React, { useState } from "react";
import useSWR, { mutate } from "swr";
import { useForm } from "react-hook-form";
import { useToasts } from "react-toast-notifications";

import Layout from "../../../components/layout";
import Button from "../../../components/button";
import { Label, Text } from "../../../components/form";
import Modal from "../../../components/modal";

import { Department } from "../../../interfaces";
import useSWRPost from "../../../hooks/useSWRPost";
import useFormErrors from "../../../hooks/useFormErrors";
import { ReactComponent as Loader } from "../../../assets/three-dots.svg";

import styles from "./style.module.css";

interface Payload {
  name: string;
}

const AddDepartment = () => {
  const [isOpen, setIsOpen] = useState<boolean>(false);
  const { addToast } = useToasts();
  const { register, handleSubmit, errors } = useForm<Payload>();
  useFormErrors(errors);

  const [createDepartment, { isValidating }] = useSWRPost<string>("/api/department", {
    onSuccess: (res) => {
      if (res.status === "success") {
        mutate("/api/departments");
        addToast(res?.message || "New department created successfully", { appearance: "success" });
        setIsOpen(false);
      } else addToast(res?.error || "Something went wrong", { appearance: "error" });
    },
    onError: (err) => {
      addToast(err?.error || "Something went wrong.", { appearance: "error" });
    },
  });

  const onSubmit = (values: Payload) => {
    createDepartment(JSON.stringify(values));
  };
  return (
    <>
      <Button
        type="button"
        style={{ background: "var(--color-brand)", width: "max-content" }}
        onClick={() => setIsOpen(true)}
      >
        Add New Department
      </Button>
      <Modal isOpen={isOpen} toggleModal={setIsOpen}>
        <form onSubmit={handleSubmit(onSubmit)} className={styles.addForm}>
          <h2>Create New Department</h2>
          <Label value="Name">
            <Text
              name="name"
              placeholder="Name of the Department"
              inpRef={register({
                required: true,
              })}
            />
          </Label>
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
      </Modal>
    </>
  );
};

const DepartmentRow: React.FC<{ department: Department }> = ({ department }) => {
  const { id, name } = department;
  return (
    <div className={styles.department}>
      <span className={styles.idSpan}>{id}</span>
      <span className={styles.nameSpan}>{name}</span>
    </div>
  );
};

const Departments = () => {
  const { data } = useSWR<{ status: string; departments: Department[] }>(`/api/departments`);

  const list = data?.departments?.map((d) => <DepartmentRow department={d} key={d.id} />);

  return (
    <Layout>
      <div className={styles.main}>
        <h1 className={styles.head}>All Departments</h1>
        <div className={styles.options}>
          <AddDepartment />
        </div>
        <div className={styles.list}>
          <div className={`${styles.department} ${styles.topDepartment}`}>
            <span className={styles.idSpan}>ID</span>
            <span className={styles.nameSpan}>Name</span>
          </div>
          {list}
        </div>
      </div>
    </Layout>
  );
};

export default Departments;
