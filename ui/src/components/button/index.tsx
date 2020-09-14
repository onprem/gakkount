import React from "react";

import styles from "./button.module.css";

export type ButtonProps = React.DetailedHTMLProps<
  React.ButtonHTMLAttributes<HTMLButtonElement>,
  HTMLButtonElement
>;

export const Button: React.FC<ButtonProps> = ({ children, ...rest }) => {
  return (
    <button {...rest} className={styles.button}>
      {children}
    </button>
  );
};

export default Button;
