import React, { useState } from "react";
import { Grid, Form, Segment, Message, Input, Button } from "semantic-ui-react";

const RegisterForm = props => {
  const [state, setState] = useState({
    name: "",
    username: "",
    email: "",
    password: "",
    confirmPass: "",
    error: ""
  });

  return (
    <Form>
      <Segment>
        <Form.Input
          icon="smile"
          iconPosition="left"
          placeholder="Enter your name"
          control={Input}
          onChange={e => {
            setState({
              ...state,
              name: e.target.value
            });
          }}
        />
        <Form.Input
          icon="user"
          iconPosition="left"
          placeholder="Enter your username"
          control={Input}
          onChange={e => {
            setState({
              ...state,
              username: e.target.value
            });
          }}
        />
        <Form.Input
          icon="mail"
          iconPosition="left"
          placeholder="Enter your email"
          control={Input}
          onChange={e => {
            setState({
              ...state,
              email: e.target.value
            });
          }}
        />
        <Form.Input
          icon="lock"
          iconPosition="left"
          type="password"
          placeholder="Choose a password"
          control={Input}
          onChange={e => {
            setState({
              ...state,
              password: e.target.value
            });
          }}
        />
        <Form.Input
          icon="lock"
          iconPosition="left"
          type="password"
          placeholder="Repeat your password"
          control={Input}
          onChange={e => {
            setState({
              ...state,
              confirmPass: e.target.value
            });
          }}
        />
        <Form.Button
          fluid
          color="blue"
          content="Register"
          control={Button}
          onClick={() => {
            //   handleRegister();
          }}
        />
      </Segment>
      <Message>
        Already have an account? <a href="/login">Login</a>
      </Message>
      {state.error !== "" && <Message negative>{state.error}</Message>}
    </Form>
  );
};

function Signup() {
  return (
    <div style={{ top: "80px", position: "absolute", width: "100%" }}>
      <Grid
        centered
        verticalAlign="middle"
        textAlign="center"
        style={{ height: "100vh" }}
      >
        <Grid.Row>
          <Grid.Column
            style={{ maxWidth: 450, marginRight: 20, marginLeft: 20 }}
          >
            <RegisterForm />
          </Grid.Column>
        </Grid.Row>
      </Grid>
    </div>
  );
}

export default Signup;
