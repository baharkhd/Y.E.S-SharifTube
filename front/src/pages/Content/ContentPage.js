import React, { useState } from "react";
import {
  Embed,
  Placeholder,
  Segment,
  Grid,
  Feed,
  Icon,
  Divider,
  Container,
  Input
} from "semantic-ui-react";
import "./ContentPage.css";
import { useParams } from "react-router-dom";

const comments = [
  {
    author: "bahar",
    body: "hii guys! :)",
    timestamp: "10 Feb",
    replies: [
      {
        author: "khodabakhshian",
        body: "hi bahar:D",
        timestamp: "11 Feb"
      },
      {
        author: "David",
        body: "bye bahar:(",
        timestamp: "12 Feb"
      },
      {
        author: "David",
        body: "bye bahar:(",
        timestamp: "12 Feb"
      },
      {
        author: "khodabakhshian",
        body: "hi bahar:D",
        timestamp: "11 Feb"
      }
    ]
  },
  {
    author: "bahar2",
    body: "hii guys! :)2",
    timestamp: "10 Feb2",
    replies: [
      {
        author: "khodabakhshian2",
        body: "hi bahar:D2",
        timestamp: "11 Feb2"
      }
    ]
  },
  {
    author: "bahar3",
    body: "hii guys! :)3",
    timestamp: "10 Feb3",
    replies: [
      {
        author: "khodabakhshian3",
        body: "hi bahar:D3",
        timestamp: "11 Feb3"
      }
    ]
  },
  {
    author: "bahar",
    body: "hii guys! :)",
    timestamp: "10 Feb",
    replies: [
      {
        author: "khodabakhshian",
        body: "hi bahar:D",
        timestamp: "11 Feb"
      },
      {
        author: "David",
        body: "bye bahar:(",
        timestamp: "12 Feb"
      },
      {
        author: "David",
        body: "bye bahar:(",
        timestamp: "12 Feb"
      },
      {
        author: "khodabakhshian",
        body: "hi bahar:D",
        timestamp: "11 Feb"
      }
    ]
  },
  {
    author: "bahar",
    body: "hii guys! :)",
    timestamp: "10 Feb",
    replies: [
      {
        author: "khodabakhshian",
        body: "hi bahar:D",
        timestamp: "11 Feb"
      },
      {
        author: "David",
        body: "bye bahar:(",
        timestamp: "12 Feb"
      },
      {
        author: "David",
        body: "bye bahar:(",
        timestamp: "12 Feb"
      },
      {
        author: "khodabakhshian",
        body: "hi bahar:D",
        timestamp: "11 Feb"
      }
    ]
  },
  {
    author: "bahar",
    body: "hii guys! :)",
    timestamp: "10 Feb",
    replies: [
      {
        author: "khodabakhshian",
        body: "hi bahar:D",
        timestamp: "11 Feb"
      },
      {
        author: "David",
        body: "bye bahar:(",
        timestamp: "12 Feb"
      },
      {
        author: "David",
        body: "bye bahar:(",
        timestamp: "12 Feb"
      },
      {
        author: "khodabakhshian",
        body: "hi bahar:D",
        timestamp: "11 Feb"
      }
    ]
  }
];

const Reply = ({ author, body, time }) => {
  return (
    <Feed.Event>
      {/* <Feed.Label>
        <img src="https://react.semantic-ui.com/images/avatar/small/elliot.jpg" />
        <Icon name="reply" />
      </Feed.Label> */}
      <Feed.Content>
        <Feed.Summary>
          <Icon name="reply" />
          <Feed.User>{author}</Feed.User> {body}
          <Feed.Date>{time}</Feed.Date>
        </Feed.Summary>
      </Feed.Content>
    </Feed.Event>
  );
};

const Comment = ({ comment }) => {
  return (
    <Feed.Event>
      <Feed.Content>
        <Feed.Summary>
          {/* <Feed.Extra>
            <Icon
              name="comment"
              color="green"
              onClick={() => {
                // Todo: add e reply to this comment
              }}
            />
          </Feed.Extra> */}
          <Feed.User>{comment.author}</Feed.User> {comment.body}
          <Feed.Date>{comment.timestamp}</Feed.Date>
        </Feed.Summary>
        <Feed.Meta>
          {comment.replies.map(reply => {
            return (
              <Reply
                author={reply.author}
                time={reply.timestamp}
                body={reply.body}
              />
            );
          })}
        </Feed.Meta>
      </Feed.Content>
    </Feed.Event>
  );
};

function ContentPage(props) {
  let { courseID, contentID } = useParams();
  courseID = courseID.substring(1);
  contentID = contentID.substring(1);
  console.log("contentID:", contentID, ", courseID:", courseID)

  const [newComment, setNewComment] = useState("");

  return (
    <div>
      <Segment style={{ top: 70, overflow: "hidden", borderRadius: 0 }}>
        <Grid columns={2} textAlign="center" fluid stackable>
          <Grid.Column>
            <Segment>
              <Placeholder className="test" fluid>
                <Placeholder.Image rectangular />
              </Placeholder>
              <Container textAlign="left">
                <h1>Title</h1>
                <p>
                  Lorem ipsum dolor sit amet, consectetuer adipiscing elit.
                  Aenean commodo ligula eget dolor. Aenean massa strong. Cum
                  sociis natoque penatibus et magnis dis parturient montes,
                  nascetur ridiculus mus. Donec quam felis, ultricies nec,
                  pellentesque eu, pretium quis, sem. Nulla consequat massa quis
                  enim. Donec pede justo, fringilla vel, aliquet nec, vulputate
                  eget, arcu. In enim justo, rhoncus ut, imperdiet a, venenatis
                  vitae, justo. Nullam dictum felis eu pede link mollis pretium.
                  Integer tincidunt. Cras dapibus. Vivamus elementum semper
                  nisi. Aenean vulputate eleifend tellus. Aenean leo ligula,
                  porttitor eu, consequat vitae, eleifend ac, enim. Aliquam
                  lorem ante, dapibus in, viverra quis, feugiat a, tellus.
                  Phasellus viverra nulla ut metus varius laoreet. Quisque
                  rutrum. Aenean imperdiet. Etiam ultricies nisi vel augue.
                  Curabitur ullamcorper ultricies nisi.
                </p>
              </Container>
            </Segment>
          </Grid.Column>
          <Grid.Column>
            <Segment>
              <Input
                fluid
                type="string"
                action={{
                  color: "blue",
                  icon: "plus"
                }}
                onChange={e => {
                  setNewComment(e.target.value);
                }}
                actionPosition="right"
                placeholder="Add a comment ..."
                // defaultValue="52.03"
              />
              <Feed>
                {comments.map(comment => {
                  return <Comment comment={comment} />;
                })}
              </Feed>
            </Segment>
          </Grid.Column>
        </Grid>
      </Segment>
    </div>
  );
}

export default ContentPage;
