# frozen_string_literal: true

class ItemsController < ApplicationController
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

  def collection
    @collection ||= ItemSearcher.search(Item.order(date: :desc), params).includes(:category)
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
