import React from "react";
import {
  Embed,
  Placeholder,
  Segment,
  Grid,
  Feed,
  Icon,
  Divider
} from "semantic-ui-react";
import "./ContentPage.css";

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
  },
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
      <Segment style={{ top: 70, overflow: "hidden", borderRadius: 0 }}>
        <Grid columns={2} textAlign="center" fluid >
          <Grid.Column>
            <Placeholder className="test" fluid>
              <Placeholder.Image rectangular />
            </Placeholder>
          </Grid.Column>
          <Grid.Column>
            <Segment>
              <Feed>
                {comments.map(comment => {
                  return (
                    <Feed.Event>
                      {/* <Feed.Label>
                              <img src="https://react.semantic-ui.com/images/avatar/small/elliot.jpg" />
                            </Feed.Label> */}
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
            </Segment>
          </Grid.Column>
          {/* <Grid.Row fluid>
            <Segment >
              <Placeholder fluid>
                <Placeholder.Image rectangular />
              </Placeholder>
              <Grid columns={1} style={{ margin: 10 }} textAlign="center">
                <Grid.Row>
                  
                </Grid.Row>
              </Grid>
            </Segment>
          </Grid.Row> */}
        </Grid>
      </Segment>
    </div>
  );
}

export default ContentPage;

// <Segment style={{ top: 70 }}>
//         <Grid columns={1} textAlign="center" fluid className="test">
//           <Grid.Row fluid>
//             <Segment >
//               <Placeholder fluid>
//                 <Placeholder.Image rectangular />
//               </Placeholder>
//               <Grid columns={1} style={{ margin: 10 }} textAlign="center">
//                 <Grid.Row>
//                   <Segment>
//                     <Feed>
//                       {comments.map(comment => {
//                         return (
//                           <Feed.Event>
//                             {/* <Feed.Label>
//                               <img src="https://react.semantic-ui.com/images/avatar/small/elliot.jpg" />
//                             </Feed.Label> */}
//                             <Feed.Content>
//                               <Feed.Summary>
//                                 <Feed.User>{comment.author}</Feed.User>{" "}
//                                 {comment.body}
//                                 <Feed.Date>{comment.timestamp}</Feed.Date>
//                               </Feed.Summary>
//                               <Feed.Meta>
//                                 {comment.replies.map(reply => {
//                                   return (
//                                     <Reply
//                                       author={reply.author}
//                                       time={reply.timestamp}
//                                       body={reply.body}
//                                     />
//                                   );
//                                 })}
//                               </Feed.Meta>
//                             </Feed.Content>
//                           </Feed.Event>
//                         );
//                       })}
//                     </Feed>
//                   </Segment>
//                 </Grid.Row>
//               </Grid>
//             </Segment>
//           </Grid.Row>
//         </Grid>
//       </Segment>
