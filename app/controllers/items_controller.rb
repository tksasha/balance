class ItemsController < ApplicationController
  include ActsAsRESTController

  helper_method :items, :consolidates

  private
  # TODO: spec me
  def collection
    items DateRange.new(DateFactory.build params).month
  end

  def items date_range
    @items ||= Item.search(date_range, params[:category]).includes(:category)
  end

  def consolidates
    @consolidates ||= Consolidate.by DateRange.new(DateFactory.build params).month
  end

  # TODO: spec me
  def resource_params
    params.require(:item).permit(:date, :formula, :category_id, :description)
  end
end
