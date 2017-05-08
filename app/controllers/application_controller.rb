class ApplicationController < ActionController::Base
  protect_from_forgery

  helper_method :collection, :resource

  before_action :build_resource, only: :create

  rescue_from ActiveRecord::RecordInvalid do
    render :errors, status: :unprocessable_entity
  end

  def new
    initialize_resource
  end

  def create
    respond_to do |format|
      format.html do
        if resource.save
          redirect_to resource
        else
          render :new
        end
      end

      format.json do
        render :errors unless resource.save
      end

      format.js do
        resource.save!
      end
    end
  end

  def update
    respond_to do |format|
      format.html do
        if resource.update resource_params
          redirect_to resource
        else
          render :edit
        end
      end

      format.json do
        render :errors unless resource.update resource_params
      end

      format.js do
        resource.update! resource_params
      end
    end
  end

  def destroy
    resource.destroy

    respond_to do |format|
      format.html { redirect_to /\A(.*)Controller\z/.match(self.class.name)[1].downcase.to_sym }

      format.js {  }
    end
  end

  private
  def resource_name
    /\A(.*)Controller\z/.match(self.class.name)[1].singularize.constantize
  end

  def collection 
    @collection ||= resource_name.all
  end

  def resource
    @resource ||= resource_name.find params[:id]
  end

  def initialize_resource
    @resource = resource_name.new
  end

  def build_resource
    @resource = resource_name.new resource_params
  end
end
