import { gql, useMutation } from "@apollo/client";
import _ from "lodash";
import React, { useState } from "react";
import { Button, Form, Input, Modal } from "semantic-ui-react";
// import { gql } from "graphql-tag";
import constants from "../../constants";


const createCoursePanelLStyle={
  backgroundColor:'#3b3261',
  fontColor:'#ffffff'
}

const createCoursePanelHFLStyle={
  backgroundColor:'#e5700a'
}

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

function AddCourseModal({ addingCourse, setState, makeNotif }) {
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

      // const data2

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
        makeNotif("Success!", "Course successfully created .", "success");
      } else {
        makeNotif("Error!", createCourse.message, "danger");
      }
    }
  });

  return (
    <Modal open={addingCourse}>
      <Modal.Header style={createCoursePanelHFLStyle}>Add a new class</Modal.Header>
      <Modal.Content scrolling style={createCoursePanelLStyle}>
        <Form inverted>
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
            style={{resize:'none', height:'180px'}}
          />
        </Form>
      </Modal.Content>
      <Modal.Actions style={createCoursePanelLStyle}>
        <Button
          positive
          onClick={() => {
            // Add class
            if (
              inputs.title.trim() !== "" &&
              inputs.summary.trim() !== "" &&
              inputs.token.trim() !== ""
            ) {
              createCourse();
            } else {
              makeNotif("Error!", constants.EMPTY_FIELDS, "danger");
            }

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
