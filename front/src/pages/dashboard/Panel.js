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
  Header
} from "semantic-ui-react";
import gql from "graphql-tag";
import { useMutation } from "@apollo/client";

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
      password
      email
    }
  }
`;

const genderOptions = [
  { key: "m", text: "Male", value: "male" },
  { key: "f", text: "Female", value: "female" },
  { key: "o", text: "Other", value: "other" }
];

const UpdatePanelModal = ({ modalOpen, setModalOpen, user }) => {
  console.log("user in update modal:", user);
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
    onCompleted: ({ updateUser }) => {
      console.log("updateUser", updateUser);
      console.log("new state:", state);
    },
    update(cache, { data: { updateUser } }) {
      const data = cache.readQuery({
        query: GET_USER_QUERY
      });

      const localData = _.cloneDeep(data);
      console.log("localData:", localData);
      // localData.user = localData.user.map(post => {
      //   return post.id === updatePost.id ? updatePost : post;
      // });

      let newName = updateUser.name == "" ? localData.user.name : updateUser.name;
      let newPassword =
        updateUser.password == "" ? localData.user.password : updateUser.password;
      let newEmail =
        updateUser.email == "" ? localData.user.email : updateUser.email;

      console.log("?????", {
        name: newName,
        password: newPassword,
        email: newEmail,
        username: localData.user.username
      });

      cache.writeQuery({
        query: GET_USER_QUERY,
        data: {
          user: {
            __typename: "User",
            name: newName,
            email: newEmail,
            username: localData.username
          }
        }
      });
    }
  });

  return (
    <Modal open={modalOpen}>
      <Modal.Header>Update your account!</Modal.Header>
      <Modal.Content>
        <Form>
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
            {/* 
            <Form.Field
              control={Select}
              options={genderOptions}
              label={{
                children: "Gender",
                htmlFor: "form-select-control-gender-update"
              }}
              placeholder="Gender"
              search
              searchInput={{ id: "form-select-control-gender-update" }}
            /> */}
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
      <Modal.Actions>
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
    <Segment
      style={{
        position: "absolute",
        left: props.isMobile ? 0 : 250,
        right: 0,
        margin: 30,
        top: 70
      }}
    >
      <Form>
        <Form.Group widths="equal">
          <Form.Field
            id="form-input-control-first-name"
            control={Input}
            label="Name"
            placeholder={props.user.name}
          />

          {/* <Form.Field
            control={Select}
            options={genderOptions}
            label={{
              children: "Gender",
              htmlFor: "form-select-control-gender"
            }}
            placeholder="Gender"
            search
            searchInput={{ id: "form-select-control-gender" }}
          /> */}
        </Form.Group>
        <Form.Group widths="equal">
          <Form.Field
            id="form-input-control-user-name"
            control={Input}
            label="User name"
            placeholder={props.user.username}
          />
          {/* <Form.Field>
            <label>Password</label>
            <Input type="password" placeholder={props.user.password} />
          </Form.Field> */}
        </Form.Group>
        <Form.Field>
          <label>Email</label>
          <Input type="email" placeholder={props.user.email} />
        </Form.Field>

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
