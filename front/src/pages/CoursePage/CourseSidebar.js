import React, { useState } from "react";
import { Sidebar, Menu, Button, Icon, List } from "semantic-ui-react";
import { useHistory, Link, useParams } from "react-router-dom";
import { gql, useMutation } from "@apollo/client";
import AddTAModal from "./AddTAModal";
import PendingPage from "./PendingPage";

const TAs = [
  "folan1",
  "folan2",
  "folan3",
  "folan4",
  "folan5",
  "folan6",
  "folan7"
];

// const ADD_TA_MUTATION = gql`
//   mutation AddTA($courseID: String!, $targetUsername: String!) {
//     promoteUserToTA(courseID: $courseID, targetUsername: $targetUsername) {
//       __typename
//       ... on Course {
//         id
//         title
//         summary
//         createdAt
//       }
//       ... on Exception {
//         message
//       }
//     }
//   }
// `;

function SideBar(props) {
  const [state, setState] = useState({
    activeItem: "Personal Information",
    addingTA: false
  });

  let isProfTA =
    props.courseTAs.some(ta => ta.username === props.username) || props.isProf;

  console.log(
    "{}{}{}",
    props.courseTAs.some(ta => ta.username === props.username),
    props.isProf,
    isProfTA
  );

  // const [promoteUserToTA] = useMutation(ADD_TA_MUTATION, {
  //   variables: {}
  // })

  let { id } = useParams();
  id = id.substring(1);

  const history = useHistory();

  const handleItemClick = (e, { name }) => setState({ activeItem: name });

  console.log("course TAs:", props.courseTAs);
  console.log("username:", props.username);
  console.log(
    "is a TA:",
    props.courseTAs.some(ta => ta.username === props.username)
  );

  let uploadPath = isProfTA
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
      <AddTAModal
        open={state.addingTA}
        setOpen={setState}
        courseID={id}
        students={props.students}
      />
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
          {props.courseTAs.map(TA => {
            return (
              <List.Item as="li">
                {/* <List.Icon name="user" /> */}
                {/* <List.Content>{TA}</List.Content> */}
                {TA.name}
              </List.Item>
            );
          })}
        </List>
      </Menu.Item>

      {isProfTA && (
        <Menu.Item>
          <Button
            positive
            onClick={() => {
              setState({ ...state, addingTA: true });
            }}
          >
            Add TA
          </Button>
        </Menu.Item>
      )}

      <Menu.Item>
        <Link to={uploadPath + "/video"}>
          <Button color="blue">Upload Videos</Button>
        </Link>
      </Menu.Item>
      {isProfTA && (
        <Menu.Item>
          <Link to={uploadPath + "/attachment"}>
            <Button color="blue">Upload Attachments</Button>
          </Link>
        </Menu.Item>
      )}
      {isProfTA && (
        // <Link to={"/course:" + id + "/pendings"} component={PendingPage} >
        <Menu.Item>
          <Link to={"/course:" + id + "/pendings"}>
            <Button color="black">Pending Contents</Button>
          </Link>
        </Menu.Item>
        // </Link>
      )}
    </Sidebar>
  );
}

export default SideBar;
