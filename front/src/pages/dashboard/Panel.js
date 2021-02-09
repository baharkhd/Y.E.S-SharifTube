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
  Menu
} from "semantic-ui-react";

const genderOptions = [
  { key: "m", text: "Male", value: "male" },
  { key: "f", text: "Female", value: "female" },
  { key: "o", text: "Other", value: "other" }
];

const FormInput = ({ label, placeholder }) => {
  return <Input label={label} placeholder={placeholder} />;
};

function Panel(props) {
  return (
    <div>
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
              id="form-input-control-last-name"
              control={Input}
              label="Last name"
              placeholder="Last name"
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
            <Button positive>Update</Button>
          </Form.Field>
        </Form>
      </Segment>
    </div>
  );
}

export default Panel;
