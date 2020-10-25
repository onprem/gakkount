import React from "react";
import useSWR from "swr";
import Layout from "../../../components/layout";

import { Course } from "../../../interfaces";
import styles from "./style.module.css";

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
  const { data, isValidating } = useSWR<{ status: string; courses: Course[] }>(`/api/courses`);

  if (isValidating) return <h2>Loading...</h2>;

  const list = data?.courses.map((c) => <CourseRow course={c} key={c.id} />);

  return (
    <Layout>
      <div className={styles.main}>
        <h1 className={styles.head}>All Courses</h1>
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
