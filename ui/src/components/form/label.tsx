import React from "react";
import styles from "./style.module.css";

export type LabelProps = { value: string } & React.DetailedHTMLProps<
  React.LabelHTMLAttributes<HTMLLabelElement>,
  HTMLLabelElement
>;
export const Label: React.FC<LabelProps> = ({ value, children, ...rest }) => {
  console.log({});
  return (
    <label {...rest} className={styles.label}>
      <span className={styles.labelSpan}>{value}</span>
      {children}
    </label>
  );
};

export default Label;
