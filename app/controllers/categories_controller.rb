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

  private

  attr_reader :resource

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
end
