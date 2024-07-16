# frozen_string_literal: true

class CategoriesController < ApplicationController
  before_action :initialize_resource, only: :new

  before_action :build_resource, only: :create

  before_action -> { response.status = :created }, only: :create

  def new; end

  def create
    respond_to do |format|
      format.js do
        render :new, status: :unprocessable_entity unless resource.save
      end
    end
  end

  def update
    respond_to do |format|
      format.js do
        render :edit, status: :unprocessable_entity unless resource.update(resource_params)
      end
    end
  end

  def destroy
    respond_to do |format|
      format.js do
        resource.destroy
      end
    end
  end

  private

  helper_method :collection, :resource

  def scope
    Category.order(:name)
  end

  def collection
    @collection ||= ::CategorySearcher.search(scope, params)
  end

  def initialize_resource
    @resource = Category.new
  end

  def resource_params
    params.require(:category).permit(:name).merge(currency: params[:currency])
  end

  def build_resource
    @resource = Category.new(resource_params)
  end

  def resource
    @resource ||= Category.find(params[:id])
  end
end
