# frozen_string_literal: true

RSpec.describe TagsController, type: :controller do
  describe '#collection' do
    context do
      before { subject.instance_variable_set :@collection, :collection }

      its(:collection) { should eq :collection }
    end

    context do
      before { allow(subject).to receive(:params).and_return(:params) }

      before { allow(Tags::GetCollectionService).to receive(:call).with(:params).and_return(:collection) }

      its(:collection) { should eq :collection }
    end
  end

  describe '#serializer' do
    let(:tag) { build :tag }

    it { expect(subject.send(:serializer, tag)).to be_a TagSerializer }
  end
end
