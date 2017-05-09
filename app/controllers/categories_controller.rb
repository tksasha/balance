class CategoriesController < ApplicationController
  private
  def resource_params
    params.require(:category).permit(:name, :visible, :income, :slug)
  end
end
