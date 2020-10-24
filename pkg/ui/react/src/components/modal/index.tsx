import React from "react";
import styles from "./style.module.css";

export interface ModalProps {
  isOpen: boolean;
  toggleModal(v: boolean): void;
}

export const Modal: React.FC<ModalProps> = ({ isOpen, toggleModal, children }) => {
  if (!isOpen) return null;

  return (
    <div className={styles.container} onClick={() => toggleModal(false)}>
      <main className={styles.main} onClick={(e) => e.stopPropagation()}>
        {children}
      </main>
    </div>
  );
};

export default Modal;
