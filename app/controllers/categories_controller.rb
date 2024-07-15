# frozen_string_literal: true

class CategoriesController < ApplicationController
  before_action :initialize_resource, only: :new

  def new; end

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
end
