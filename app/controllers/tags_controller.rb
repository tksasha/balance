# frozen_string_literal: true

class TagsController < ApplicationController
  private

  def collection
    @collection ||= ::Tags::GetCollectionService.call(params)
  end

  def serializer(tag)
    TagSerializer.new(tag)
  end
end
