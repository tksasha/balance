# frozen_string_literal: true

RSpec.shared_examples :update do |params|
  before { @format = (params && params[:format]) || :js }

  describe "#update.#{ @format }" do
    let(:resource) { double }

    before { expect(subject).to receive(:resource).and_return(resource) }

    before { expect(subject).to receive(:resource_params).and_return(:params) }

    context do
      before { expect(resource).to receive(:update).with(:params).and_return(true) }

      before { patch :update, params: { id: 1 }, format: @format }

      it { success.call }
    end

    context do
      before { expect(resource).to receive(:update).with(:params).and_return(false) }

      before { patch :update, params: { id: 1 }, format: @format }

      it { failure.call }
    end
  end
end

RSpec.shared_examples :destroy do |params|
  before { @format = (params && params[:format]) || :js }

  describe "#destroy.#{ @format }" do
    let(:resource) { double }

    before { expect(subject).to receive(:resource).and_return(resource) }

    before { expect(resource).to receive(:destroy) }

    before { delete :destroy, params: { id: 1 }, format: @format }

    it { success.call }
  end
end

RSpec.shared_examples :create do |params|
  before { @format = (params && params[:format]) || :js }

  describe "#update.#{ @format }" do
    let(:resource) { double }

    before { expect(subject).to receive(:build_resource).and_return(resource) }

    before { expect(subject).to receive(:resource).and_return(resource) }

    context do
      before { expect(resource).to receive(:save).and_return(true) }

      before { post :create, format: @format }

      it { success.call }
    end

    context do
      before { expect(resource).to receive(:save).and_return(false) }

      before { patch :create, format: @format }

      it { failure.call }
    end
  end
end
