import React, { useState } from "react";
import useSWR, { mutate } from "swr";
import { useForm } from "react-hook-form";
import { useToasts } from "react-toast-notifications";

import Layout from "../../../components/layout";
import Button from "../../../components/button";
import { Label, Text } from "../../../components/form";
import Modal from "../../../components/modal";

import { Course } from "../../../interfaces";
import useSWRPost from "../../../hooks/useSWRPost";
import useFormErrors from "../../../hooks/useFormErrors";
import { ReactComponent as Loader } from "../../../assets/three-dots.svg";

import styles from "./style.module.css";

interface Payload {
  name: string;
  code: string;
  semesters: string;
}

const AddCourse = () => {
  const [isOpen, setIsOpen] = useState<boolean>(false);
  const { addToast } = useToasts();
  const { register, handleSubmit, errors } = useForm<Payload>();
  useFormErrors(errors);

  const [createCourse, { isValidating }] = useSWRPost<string>("/api/course", {
    onSuccess: (res) => {
      if (res.status === "success") {
        mutate("/api/courses");
        addToast(res?.message || "New course created successfully", { appearance: "success" });
        setIsOpen(false);
      } else addToast(res?.error || "Something went wrong", { appearance: "error" });
    },
    onError: (err) => {
      addToast(err?.error || "Something went wrong.", { appearance: "error" });
    },
  });

  const onSubmit = (values: Payload) => {
    const p = {
      ...values,
      semesters: Number(values.semesters),
    };
    createCourse(JSON.stringify(p));
  };
  return (
    <>
      <Button
        type="button"
        style={{ background: "var(--color-brand)", width: "max-content" }}
        onClick={() => setIsOpen(true)}
      >
        Add New Course
      </Button>
      <Modal isOpen={isOpen} toggleModal={setIsOpen}>
        <form onSubmit={handleSubmit(onSubmit)} className={styles.addForm}>
          <h2>Create New Course</h2>
          <Label value="Name">
            <Text
              name="name"
              placeholder="Name of the Course"
              inpRef={register({
                required: true,
              })}
            />
          </Label>
          <Label value="Code">
            <Text
              name="code"
              placeholder="Short code; ex. IMT"
              inpRef={register({
                required: true,
              })}
            />
          </Label>
          <Label value="Semesters">
            <Text
              name="semesters"
              type="number"
              placeholder="No. of semesters"
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

const CourseRow: React.FC<{ course: Course }> = ({ course }) => {
  const { id, name, code, semesters } = course;
  return (
    <div className={styles.course}>
      <span className={styles.idSpan}>{id}</span>
      <span className={styles.nameSpan}>{name}</span>
      <span className={styles.codeSpan}>{code}</span>
      <span className={styles.semSpan}>{semesters}</span>
    </div>
  );
};

const Courses = () => {
  const { data } = useSWR<{ status: string; courses: Course[] }>(`/api/courses`);

  const list = data?.courses.map((c) => <CourseRow course={c} key={c.id} />);

  return (
    <Layout>
      <div className={styles.main}>
        <h1 className={styles.head}>All Courses</h1>
        <div className={styles.options}>
          <AddCourse />
        </div>
        <div className={styles.list}>
          <div className={`${styles.course} ${styles.topCourse}`}>
            <span className={styles.idSpan}>ID</span>
            <span className={styles.nameSpan}>Name</span>
            <span className={styles.codeSpan}>Code</span>
            <span className={styles.semSpan}>Semesters</span>
          </div>
          {list}
        </div>
      </div>
    </Layout>
  );
};

export default Courses;
