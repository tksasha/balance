RSpec.shared_examples :edit do |params|
  before { @format = (params && params[:format]) || :js }

  describe "#edit.#{ @format }" do
    before { get :edit, params: { id: 1 }, xhr: (@format == :js), format: @format }

    it { should render_template :edit }
  end
end

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
