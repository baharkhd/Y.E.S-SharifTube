import React from "react";
import {
  Embed,
  Placeholder,
  Segment,
  Grid,
  Feed,
  Icon
} from "semantic-ui-react";

// type Comment{
//     id: ID!
//     author: User!
//     body: String!
//     timestamp: Int!
//     replies: [Reply!]
//     content: Content!
// }

// type Reply{
//     id: ID!
//     author: User!
//     body: String!
//     timestamp: Int!
//     comment: Comment!
// }

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

function ContentPage(props) {
  return (
    <div>
      <Segment style={{ top: 70 }}>
        <Placeholder>
          <Placeholder.Image rectangular />
        </Placeholder>
        <Grid columns={1} style={{ margin: 10 }}>
          <Grid.Row>
            <Feed>
              {comments.map(comment => {
                return (
                  <Feed.Event>
                    <Feed.Label>
                      <img src="https://react.semantic-ui.com/images/avatar/small/elliot.jpg" />
                    </Feed.Label>
                    <Feed.Content>
                      <Feed.Summary>
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
              })}
            </Feed>
          </Grid.Row>
        </Grid>
      </Segment>
    </div>
  );
}

export default ContentPage;
