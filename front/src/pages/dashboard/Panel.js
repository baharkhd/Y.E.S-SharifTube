import React, { useState } from "react";
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

const genderOptions = [
  { key: "m", text: "Male", value: "male" },
  { key: "f", text: "Female", value: "female" },
  { key: "o", text: "Other", value: "other" }
];

const UpdatePanelModal = ({ modalOpen, setModalOpen }) => {
  console.log("open:", modalOpen);
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
              placeholder="First name"
            />

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
            />
          </Form.Group>
          <Form.Group widths="equal">
            <Form.Field
              id="form-input-control-user-name-update"
              control={Input}
              label="User name"
              placeholder="User name"
            />
            <Form.Field>
              <label>Password</label>
              <Input type="password" />
            </Form.Field>
          </Form.Group>
          <Form.Field>
            <label>Email</label>
            <Input type="email" placeholder="Email" />
          </Form.Field>
        </Form>
      </Modal.Content>
      <Modal.Actions>
        <Button
          positive
          primary
          onClick={() => {
            // Todo: update information
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
            placeholder="First name"
          />

          <Form.Field
            control={Select}
            options={genderOptions}
            label={{
              children: "Gender",
              htmlFor: "form-select-control-gender"
            }}
            placeholder="Gender"
            search
            searchInput={{ id: "form-select-control-gender" }}
          />
        </Form.Group>
        <Form.Group widths="equal">
          <Form.Field
            id="form-input-control-user-name"
            control={Input}
            label="User name"
            placeholder="User name"
          />
          <Form.Field>
            <label>Password</label>
            <Input type="password" />
          </Form.Field>
        </Form.Group>
        <Form.Field>
          <label>Email</label>
          <Input type="email" placeholder="Email" />
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
  const [modalOpen, setModalOpen] = useState(true);
  return (
    <div>
      <UpdatePanelModal modalOpen={modalOpen} setModalOpen={setModalOpen} />
      <PanelInfo setModalOpen={setModalOpen} />
    </div>
  );
}

export default Panel;
