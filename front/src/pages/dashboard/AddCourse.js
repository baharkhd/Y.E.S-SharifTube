import React, { useState } from "react";
import { Modal, Button, Form, Label, Input, TextArea } from "semantic-ui-react";
import { useMutation, gql } from "@apollo/client";
import _ from "lodash";
// import { gql } from "graphql-tag";

const CREATE_COURSE_MUTATION = gql`
  mutation CreateCourse($title: String!, $summary: String, $token: String!) {
    createCourse(target: { title: $title, summary: $summary, token: $token }) {
      __typename
      ... on Course {
        id
        title
        summary
        createdAt
      }
      __typename
      ... on Exception {
        message
      }
    }
  }
`;

const COURSES_QUERY = gql`
  query GetCoursesByFilter($keyWords: [String!]!, $amount: Int!, $start: Int!) {
    coursesByKeyWords(keyWords: $keyWords, amount: $amount, start: $start) {
      id
      title
      summary
      createdAt
      # prof {
      #   username
      #   name
      # }
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
    update(cache, { data: { createCourse } }) {
      const data = cache.readQuery({
        query: COURSES_QUERY,
        variables: {
          keyWords: [],
          amount: 100,
          start: 0
        }
      });

      const localData = _.cloneDeep(data);
      console.log("course added in add course:", createCourse);
      console.log("local data in add course:", localData);

      localData.coursesByKeyWords = [
        ...localData.coursesByKeyWords,
        createCourse
      ];

      console.log("local data after chaning?:", localData);

      cache.writeQuery({
        query: COURSES_QUERY,
        data: {
          ...localData
        }
      });
    },
    onCompleted: ({ createCourse }) => {
      console.log("createCourse:", createCourse);
      if (createCourse.__typename == "Course") {
        alert("you successfully created your own class :D");
      } else {
        alert(createCourse.message);
      }
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
