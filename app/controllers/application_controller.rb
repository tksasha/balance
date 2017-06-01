class ApplicationController < ActionController::Base
  protect_from_forgery

  helper_method :collection, :resource

  before_action :initialize_resource, only: :new

  before_action :build_resource, only: :create

  rescue_from ActiveRecord::RecordInvalid do
    render :errors, status: :unprocessable_entity
  end

  def create
    if resource.save
      render :create
    else
      render :errors
    end
  end

  def update
    resource.update! resource_params
  end

  def destroy
    resource.destroy
  end
end
