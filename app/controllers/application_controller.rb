class ApplicationController < ActionController::Base
  protect_from_forgery

  before_action :initialize_resource, only: :new

  before_action :build_resource, only: :create

  before_action -> { response.status = 201 }, only: :create

  helper_method :collection, :resource
end
