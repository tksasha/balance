class CategoriesController < ApplicationController
  include ActsAsRESTController

  private
  def collection
    @collection ||= Category.order :income
  end

  def resource_params
    params.require(:category).permit(:name, :income)
  end
end
