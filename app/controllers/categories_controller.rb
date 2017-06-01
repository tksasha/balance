class CategoriesController < ApplicationController
  include ActsAsRESTController

  private
  def collection
    @collection ||= Category.order :income
  end
end
