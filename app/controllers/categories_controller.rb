class CategoriesController < ApplicationController
  include ActsAsRESTController

  private
  def collection
    @collection ||= Category.order :income, :name
  end
end
