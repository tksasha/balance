# frozen_string_literal: true

class CashesController < ApplicationController
  def update
    render :edit, status: :unprocessable_entity unless resource.update(resource_params)
  end

  private

  def collection
    @collection ||= ::Cashes::GetCollectionService.call(params)
  end

  def resource
    @resource ||= ::Cashes::GetResourceService.call(params)
  end

  def resource_params
    params.require(:cash).permit(:formula, :name, :currency)
  end
end
