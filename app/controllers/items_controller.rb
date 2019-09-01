# frozen_string_literal: true

class ItemsController < ApplicationController
  helper_method :items, :consolidates

  def create
    render :new, status: 422 unless resource.save
  end

  def update
    render :edit, status: 422 unless resource.update resource_params
  end

  def destroy
    resource.destroy
  end

  private

  # TODO: spec me
  def collection
    items DateRange.new(DateFactory.build(params)).month
  end

  def items(date_range)
    @items ||= Item.search(date_range, params[:category]).includes(:category)
  end

  # TODO: spec me
  def consolidates
    @consolidates ||= Consolidate.by DateRange.new(DateFactory.build(params)).month
  end

  def resource_params
    params.require(:item).permit(:date, :formula, :category_id, :description)
  end

  def resource
    @resource ||= Item.find params[:id]
  end

  def build_resource
    @resource = Item.new resource_params
  end
end
