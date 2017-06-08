module ActsAsRESTController
  extend ActiveSupport::Concern

  included do
    # TODO: spec me
    before_action :initialize_resource, only: :new

    # TODO: spec me
    before_action :build_resource, only: :create

    # TODO: spec me
    helper_method :collection, :resource
  end

  def create
    render :new unless resource.save
  end

  def update
    render :edit unless resource.update resource_params
  end

  def destroy
    resource.destroy
  end

  private
  def resource_class
    self.class.name.gsub(/Controller\z/, '').singularize.constantize
  end

  def initialize_resource
    @resource = resource_class.new
  end

  def build_resource
    @resource = resource_class.new resource_params
  end

  def resource
    @resource ||= resource_class.find params[:id]
  end
end
