class ApplicationController < ActionController::Base
  protect_from_forgery

  helper_method :collection, :resource

  before_action :build_resource, only: :create

  before_action :initialize_resource, only: :new

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
        render :errors unless resource.save
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
        render :errors unless resource.update resource_params
      end
    end
  end

  def destroy
    resource.destroy

    respond_to do |format|
      format.html do
        redirect_to resource_collection_sym
      end

      format.json {}

      format.js {}
    end
  end

  private

  #
  # CategoriesController => 'Categories'
  #
  def resource_collection_name
    @resource_collection_name ||= /\A(.*)Controller\z/.match(self.class.name)[1]
  end

  #
  # CategoriesController => :categories
  #
  def resource_collection_sym
    @resource_collection_sym ||= resource_collection_name.downcase.to_sym
  end

  #
  # CategoriesController => Category
  #
  def resource_model
    @resource_model ||= resource_collection_name.singularize.constantize
  end

  def collection 
    @collection ||= resource_model.all
  end

  def resource
    @resource ||= resource_model.find params[:id]
  end

  def initialize_resource
    @resource = resource_model.new
  end

  def build_resource
    @resource = resource_model.new resource_params
  end
end
