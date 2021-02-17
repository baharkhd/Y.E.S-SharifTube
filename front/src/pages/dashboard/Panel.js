import React, { useState } from "react";
import _ from "lodash";
import {
  Grid,
  Input,
  Segment,
  Form,
  Button,
  Checkbox,
  TextArea,
  Select,
  Menu,
  Image,
  Modal,
  Icon,
  List,
  Header, Message, Label
} from "semantic-ui-react";
import gql from "graphql-tag";
import { useMutation } from "@apollo/client";

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

const genderOptions = [
  { key: "m", text: "Male", value: "male" },
  { key: "f", text: "Female", value: "female" },
  { key: "o", text: "Other", value: "other" }
];

const UpdatePanelModal = ({ modalOpen, setModalOpen, user, setUser }) => {
  // console.log("user in update modal:", user);
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

      console.log("updateUser:", updateUser)
      console.log("????????????????????????", data);

      const localData = _.cloneDeep(data);
      console.log("localData:", localData);
      // localData.user = localData.user.map(post => {
      //   return post.id === updatePost.id ? updatePost : post;
      // });

      let newName =
        updateUser.name == "" ? localData.user.name : updateUser.name;
      let newPassword =
        updateUser.password == ""
          ? localData.user.password
          : updateUser.password;
      let newEmail =
        updateUser.email == "" ? localData.user.email : updateUser.email;

      console.log("?????", {
        name: newName,
        password: newPassword,
        email: newEmail,
        username: localData.user.username
      });

      console.log("???????????????????//", {
        __typename: "User",
        name: newName,
        email: newEmail,
        username: localData.username,
        password: ""
      });

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
      console.log("updateUser", updateUser);
      // setUser({ user: { ...user, ...updateUser } });
      console.log("new state:", state);
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
            // Todo: update information
            updateUser();
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
        margin: '30vh',
        top: '-10vh',
        borderColor: "#012968",
      }}
    >
      <Form>
        <List divided selection style={panelFormLStyle}>
          <List.Item>
            <Label color='red' Name>
              Name
            </Label>
            <span>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;{props.user.name}</span>
          </List.Item>
          <List.Item>
            <Label color='blue' Username>
              Username
            </Label>
            <span>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;{props.user.username}</span>
          </List.Item>
          <List.Item>
            <Label color='orange' Email>
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
