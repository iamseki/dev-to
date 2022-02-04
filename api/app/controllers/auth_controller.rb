class AuthController < ApplicationController
    def sign_in
        #auth = request.request_parameters[:auth]
        key = params[:key]
        # fake sign_in code, by default HTTP STATUS CODE will be 204
        render json: @controller.to_json, status: 401 if key != "supersecret"   
    end
end