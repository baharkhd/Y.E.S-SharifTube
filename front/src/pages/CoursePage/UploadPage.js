import React from "react";
import { Segment, Input, Form, TextArea } from "semantic-ui-react";

function UploadPage() {
  return (
    <Segment style={{ top: 70 }}>
      <Segment>Where you should upload videos</Segment>
      <Form>
        <Form.Group widths="four">
          <Form.Field
            control={Input}
            label="Title of this content"
            placeholder="Title"
          />
        </Form.Group>
        <Form.Field
          control={TextArea}
          label="Description of this content"
          placeholder="Write a summary about this content"
        />
      </Form>
    </Segment>
  );
}

export default UploadPage;
