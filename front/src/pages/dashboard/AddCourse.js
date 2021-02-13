import React, { useState } from "react";
import { Modal, Button, Form, Label, Input, TextArea } from "semantic-ui-react";
import { useMutation, gql } from "@apollo/client";
// import { gql } from "graphql-tag";

const CREATE_COURSE_MUTATION = gql`
  mutation CreateCourse($title: String!, $summary: String, $token: String!) {
    createCourse(target: { title: $title, summary: $summary, token: $token }) {
      __typename
      ... on Course {
        title
      }
      __typename
      ... on Exception {
        message
      }
    }
  }
`;

function AddCourseModal({ addingCourse, setState }) {
  const [inputs, setInputs] = useState({
    title: "",
    summary: "",
    token: ""
  });

  const [createCourse] = useMutation(CREATE_COURSE_MUTATION, {
    variables: inputs,
    onCompleted: ({ createCourse }) => {
      console.log("createCourse:", createCourse);
    }
  });

  return (
    <Modal open={addingCourse}>
      <Modal.Header>Add a new class</Modal.Header>
      <Modal.Content scrolling>
        <Form>
          <Form.Group widths="equal">
            <Form.Field
              control={Input}
              label="Title"
              placeholder="Enter title of your class"
              onChange={e => {
                setInputs({ ...inputs, title: e.target.value });
              }}
            />
            <Form.Field
              control={Input}
              label="Password for this class"
              placeholder="Enter password for your class"
              onChange={e => {
                setInputs({ ...inputs, token: e.target.value });
              }}
            />
          </Form.Group>
          <Form.TextArea
            label="Description"
            placeholder="Enter description of your class"
            onChange={e => {
              setInputs({ ...inputs, summary: e.target.value });
            }}
          />
        </Form>
      </Modal.Content>
      <Modal.Actions>
        <Button
          positive
          onClick={() => {
            // Add class
            createCourse();
            setState({ addingCourse: false });
          }}
        >
          Add
        </Button>
        <Button
          negative
          onClick={() => {
            setState({ addingCourse: false });
          }}
        >
          Cancel
        </Button>
      </Modal.Actions>
    </Modal>
  );
}

export default AddCourseModal;
