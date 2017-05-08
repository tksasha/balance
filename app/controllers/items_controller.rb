class ItemsController < ApplicationController
  helper_method :items, :consolidates

  #
  # TODO: add push state (change URL by ajax)
  #

  def create
    render :new unless resource.save
  end

  #
  # TODO: implement errors.js.erb for update.js
  #

  private
  def collection
    items DateRange.new(DateFactory.build params).month
  end

  def items date_range
    @items ||= Item.search(date_range, params[:category]).includes(:category)
  end

  def consolidates
    @consolidates ||= Consolidate.by DateRange.new(DateFactory.build params).month
  end

  def resource_params
    params.require(:item).permit(:date, :formula, :category_id, :description)
  end
end
