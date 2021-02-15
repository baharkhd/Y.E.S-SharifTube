import React, { useState } from "react";
import { Sidebar, Menu, Button, Icon, List } from "semantic-ui-react";
import { useHistory, Link, useParams } from "react-router-dom";
import { gql, useMutation } from "@apollo/client";
import AddTAModal from "./AddTAModal";

const TAs = [
  "folan1",
  "folan2",
  "folan3",
  "folan4",
  "folan5",
  "folan6",
  "folan7"
];

const ADD_TA_MUTATION = gql`
  mutation AddTA($courseID: String!, $targetUsername: String!) {
    promoteUserToTA(courseID: $courseID, targetUsername: $targetUsername) {
      __typename
      ... on Course {
        id
        title
        summary
        createdAt
      }
      ... on Exception {
        message
      }
    }
  }
`;

function SideBar(props) {
  const [state, setState] = useState({
    activeItem: "Personal Information",
    addingTA: false
  });

  // const [promoteUserToTA] = useMutation(ADD_TA_MUTATION, {
  //   variables: {}
  // })

  let { id } = useParams();
  id = id.substring(1);

  const history = useHistory();

  const handleItemClick = (e, { name }) => setState({ activeItem: name });

  let uploadPath =
    props.role == "prof"
      ? "/course:" + id + "/upload"
      : "/course:" + id + "/offer";

  return (
    <Sidebar
      as={Menu}
      animation="overlay"
      icon="labeled"
      direction="left"
      vertical
      visible={props.isMobile ? props.sidebarIsOpen : true}
      width="thin"
      style={{ width: 250, top: 70 }}
    >
      <AddTAModal open={state.addingTA} setOpen={setState} courseID={id} />
      <Menu.Item as="a">
        <Icon name="student" />
        Course Title: {props.courseTitle ? props.courseTitle : ""}
      </Menu.Item>
      <Menu.Item as="a">
        <Icon name="student" />
        Course Instructor: {props.courseProf.name ? props.courseProf.name : ""}
      </Menu.Item>
      <Menu.Item as="a">
        <Icon name="users" />
        TAs:
        <List>
          {TAs !== null &&
            TAs.map(TA => {
              return (
                <List.Item as="li">
                  {/* <List.Icon name="user" /> */}
                  {/* <List.Content>{TA}</List.Content> */}
                  {TA.name}
                </List.Item>
              );
            })}
          <List.Item>
            <Button
              positive
              onClick={() => {
                setState({ ...state, addingTA: true });
              }}
            >
              Add TA
            </Button>
          </List.Item>
        </List>
      </Menu.Item>
      <Menu.Item>
        <Link to={uploadPath}>
          <Button color="blue">Upload Videos</Button>
        </Link>
        {/* If user is the instructor or a TA */}
        {/* {props.role === "prof" ? (
          <Link to={"/course:" + id + "/pendings"}>
            <Button color="blue">Pending Contents</Button>
          </Link>
        ) : (
          <Link to={"/course:" + id + "/upload"}>
            <Button color="blue">Upload Videos</Button>
          </Link>
        )} */}
      </Menu.Item>
      {props.role === "prof" && (
        <Menu.Item>
          <Link to={"/course:" + id + "/pendings"}>
            <Button color="black">Pending Contents</Button>
          </Link>
        </Menu.Item>
      )}
    </Sidebar>
  );
}

export default SideBar;
