class CategoriesController < ApplicationController
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
end
