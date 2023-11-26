package handlers
import(

)

func apiResponse(status int, body interface{})(*events.APIGatewayProxyResponse,error){
	resp:=events.APIGatewayProxyResponse{
		Headers: map[string]string["Content-Type":"application/json"]}
		resp.statusCode=status
		stringBody,_:=json.Marshal(body)	
		
}