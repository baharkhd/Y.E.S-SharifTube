import React from "react";
import { Modal, Button, Form, Label, Input, TextArea } from "semantic-ui-react";

function AddCourseModal({ addingPost, setState }) {
  return (
    <Modal open={addingPost}>
      <Modal.Header>Add a new class</Modal.Header>
      <Modal.Content>
        <Form>
          <Form.Group>
            <Form.Field
              control={Input}
              label="Title"
              placeholder="Enter title of your class"
            />
          </Form.Group>
          <Form.TextArea
            label="Description"
            placeholder="Enter description of your class"
          />
        </Form>
      </Modal.Content>
      <Modal.Actions>
        <Button
          positive
          onClick={() => {
            // Add class
          }}
        >
          Add
        </Button>
        <Button
          negative
          onClick={() => {
            setState({ addingPost: false });
          }}
        >
          {" "}
          Cancel
        </Button>
      </Modal.Actions>
    </Modal>
  );
}

export default AddCourseModal;
