# frozen_string_literal: true

class ItemsController < ApplicationController
  delegate :destroy, to: :resource

  def create
    render :new, status: :unprocessable_entity unless resource.save
  end

  def update
    render :edit, status: :unprocessable_entity unless resource.update resource_params
  end

  private

  def relation
    Item.order(date: :desc).includes(:category)
  end

  def collection
    @collection ||= ItemSearcher.search relation, params
  end

  def resource_params
    params.require(:item).permit(:date, :formula, :category_id, :description, :currency)
  end

  def resource
    @resource ||= Item.find params[:id]
  end

  def build_resource
    @resource = Item.new resource_params
  end
end
