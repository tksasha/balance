# frozen_string_literal: true

class CategoriesController < BaseController
  def new; end

  private

  def scope
    Category.order(:name)
  end

  def collection
    @collection ||= ::CategorySearcher.search(scope, params)
  end
end
