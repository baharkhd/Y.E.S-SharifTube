import { useMutation } from "@apollo/client";
import gql from "graphql-tag";
import _ from "lodash";
import React, { useState } from "react";
import { Button, Form, Input, Label, List, Modal, Segment } from "semantic-ui-react";
import constants from '../../constants';

const panelFormLStyle={
  fontSize:'15px',
  overFlow:'hidden'
}

const updateUserPanelLStyle={
  backgroundColor:'#424571',
  fontColor:'#ffffff'
}

const updateUserPanelHFLStyle={
  backgroundColor:'#fffb00'
}

const UPDATE_USER_MUTATION = gql`
  mutation UpdateUser($name: String, $password: String, $email: String) {
    updateUser(toBe: { name: $name, password: $password, email: $email }) {
      __typename
      ... on User {
        name
        email
      }
    }
  }
`;

const GET_USER_QUERY = gql`
  {
    user {
      username
      name
      # password
      email
    }
  }
`;

const UpdatePanelModal = ({
  modalOpen,
  setModalOpen,
  user,
  setUser,
  makeNotif
}) => {
  const [state, setState] = useState({
    newName: user.name,
    newGender: "",
    newUsername: user.username,
    newPass: user.password,
    newEmail: user.email
  });

  const [updateUser] = useMutation(UPDATE_USER_MUTATION, {
    variables: {
      name: state.newName,
      password: state.newPass,
      email: state.newEmail
    },
    update(cache, { data: { updateUser } }) {
      const data = cache.readQuery({
        query: GET_USER_QUERY
      });

      const localData = _.cloneDeep(data);

      let newName =
        updateUser.name == "" ? localData.user.name : updateUser.name;
      let newPassword =
        updateUser.password == ""
          ? localData.user.password
          : updateUser.password;
      let newEmail =
        updateUser.email == "" ? localData.user.email : updateUser.email;


      cache.writeQuery({
        query: GET_USER_QUERY,
        data: {
          user: {
            __typename: "User",
            name: newName,
            email: newEmail,
            username: localData.username,
            password: ""
          }
        }
      });
    },
    onCompleted: ({ updateUser }) => {
      makeNotif(
        "Success!",
        "Your personal information successfully updated!",
        "success"
      );
    }
  });

  return (
    <Modal open={modalOpen}>
      <Modal.Header style={updateUserPanelHFLStyle}>Update your account!</Modal.Header>
      <Modal.Content style={updateUserPanelLStyle}>
        <Form inverted>
          <Form.Group widths="equal">
            <Form.Field
              id="form-input-control-first-name-update"
              control={Input}
              label="Name"
              placeholder="Enter new name"
              value={state.newName}
              onChange={e => {
                setState({ ...state, newName: e.target.value });
              }}
            />
          </Form.Group>
          <Form.Group widths="equal">
            <Form.Field
              id="form-input-control-user-name-update"
              control={Input}
              label="Username"
              placeholder="Username"
              value={state.newUsername}
              readOnly
            />
            <Form.Field>
              <label>Password</label>
              <Input
                type="string"
                value={state.newPass}
                placeholder="Enter new password"
                onChange={e => {
                  setState({ ...state, newPass: e.target.value });
                }}
              />
            </Form.Field>
          </Form.Group>
          <Form.Field>
            <label>Email</label>
            <Input
              type="email"
              placeholder="Enter new email"
              value={state.newEmail}
              onChange={e => {
                setState({ ...state, newEmail: e.target.value });
              }}
            />
          </Form.Field>
        </Form>
      </Modal.Content>
      <Modal.Actions style={updateUserPanelLStyle}>
        <Button
          positive
          primary
          onClick={() => {
            if (
              state.newName.trim() !== "" &&
              state.newEmail.trim() !== "" &&
              state.newPass.trim() !== ""
            ) {
              updateUser();
            } else {
              makeNotif("Error!", constants.EMPTY_FIELDS, "danger");
            }
            setModalOpen(false);
          }}
        >
          Update
        </Button>
        <Button negative onClick={() => setModalOpen(false)} primary>
          Cancel
        </Button>
      </Modal.Actions>
    </Modal>
  );
};

const PanelInfo = props => {
  return (
    <Segment raised
      style={{
        position: "absolute",
        left: props.isMobile ? 0 : 250,
        right: 0,
        margin: 45,
        top: 90,
        borderColor: "#012968",
        backgroundColor: "#cadbeb",
        boxShadow: "0 8px 6px -6px black",
      }}
    >
      <Form>
        <List divided selection style={panelFormLStyle}>
          <List.Item>
            <Label style={{backgroundColor: "#203e59", color: "white"}} Name>
              Name
            </Label>
            <span>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;{props.user.name}</span>
          </List.Item>
          <List.Item>
            <Label  Username style={{backgroundColor: "#203e59", color: "white"}}>
              Username
            </Label>
            <span>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;{props.user.username}</span>
          </List.Item>
          <List.Item>
            <Label style={{backgroundColor: "#203e59", color: "white"}} Email>
              Email
            </Label>
            <span>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;{props.user.email}</span>
          </List.Item>
        </List>
        <Form.Field>
          <Button
            positive
            onClick={() => {
              props.setModalOpen(true);
            }}
          >
            Update Info
          </Button>
        </Form.Field>
      </Form>
    </Segment>
  );
};

function Panel(props) {
  const [modalOpen, setModalOpen] = useState(false);
  return (
    <div>
      <UpdatePanelModal
        modalOpen={modalOpen}
        setModalOpen={setModalOpen}
        user={props.user}
        setUser={props.setState}
        makeNotif={props.makeNotif}
      />
      <PanelInfo
        setModalOpen={setModalOpen}
        isMobile={props.isMobile}
        user={props.user}
      />
    </div>
  );
}

export default Panel;
