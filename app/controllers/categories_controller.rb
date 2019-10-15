# frozen_string_literal: true

class CategoriesController < ApplicationController
  before_action :set_variant, only: :index

  def create
    render :new, status: 422 unless resource.save
  end

  def update
    render :edit, status: 422 unless resource.update resource_params
  end

  private

  def collection
    @collection ||= Category.order :income
  end

  def resource_params
    params.require(:category).permit(:name, :income, :visible)
  end

  def resource
    @resource ||= Category.find params[:id]
  end

  def set_variant
    request.variant = :widget if params[:widget].present?
  end

  def initialize_resource
    @resource = Category.new
  end

  def build_resource
    @resource = CategoryService.new resource_params
  end
end
